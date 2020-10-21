package tools

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
	"strings"
)

type KafkaSender struct {
	Uri      string
	UserName string
	Password string
}

type KafkaSendMsg struct {
	BrokerList  string
	Topic       string
	Value       string //要发送的消息文本
	UserName    string
	Password    string
	Key         string
	Partitioner string
	Partition   int
}

func (s *KafkaSender) KafkaCreateTopic(topic string) error {
	msg := &KafkaSendMsg{
		BrokerList: s.Uri,
		Topic:      topic,
	}

	broker := sarama.NewBroker(msg.BrokerList)
	config := sarama.NewConfig()
	config.Version = sarama.V2_5_0_0
	config.Net.SASL.Enable = true
	config.Net.SASL.User = "admin"
	config.Net.SASL.Password = "admin001"
	err := broker.Open(config)

	if err != nil {
		panic(err)
	}
	admin, err := sarama.NewClusterAdmin([]string{broker.Addr()}, config)
	if err != nil {
		return err
	}
	//err = admin.CreateTopic(topic, &sarama.TopicDetail{NumPartitions: 1, ReplicationFactor: 1}, false)
	//if err != nil {
	//	return err
	//}
	r := sarama.Resource{ResourceType: sarama.AclResourceTopic, ResourceName: topic}
	a := sarama.Acl{Principal: "User:reader01", Host: "*", Operation: sarama.AclOperationRead, PermissionType: sarama.AclPermissionAllow}
	err = admin.CreateACL(r, a)
	if err != nil {
		panic(err)
	}
	a.Principal = "User:writer01"
	a.Host = "*"
	a.Operation = sarama.AclOperationAll
	a.PermissionType = sarama.AclPermissionAllow
	err = admin.CreateACL(r, a)
	//resourceName := topic
	//filter := sarama.AclFilter{
	//	ResourceType: sarama.AclResourceTopic,
	//	Operation:    sarama.AclOperationAlter,
	//	ResourceName: &resourceName,
	//}
	//_, err = admin.DeleteACL(filter, false)
	if err != nil {
		panic(err)
	}
	defer admin.Close()
	return nil
}

func (s *KafkaSender) KafkaSend(topic string, params map[string]interface{}) error {
	data, err := json.Marshal(params)
	msg := &KafkaSendMsg{
		BrokerList: s.Uri,
		Topic:      topic,
		Value:      string(data),
		UserName:   s.UserName,
		Password:   s.Password,
	}

	if len(msg.BrokerList) == 0 {
		return errors.New("no -brokers specified. Alternatively, set the KAFKA_PEERS environment variable")
	}

	if msg.Topic == "" {
		return errors.New("no -topic specified")
	}

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	switch msg.Partitioner {
	case "":
		if msg.Partition >= 0 {
			config.Producer.Partitioner = sarama.NewManualPartitioner
		} else {
			config.Producer.Partitioner = sarama.NewHashPartitioner
		}
	case "hash":
		config.Producer.Partitioner = sarama.NewHashPartitioner
	case "random":
		config.Producer.Partitioner = sarama.NewRandomPartitioner
	case "manual":
		config.Producer.Partitioner = sarama.NewManualPartitioner
		if msg.Partition == -1 {
			return errors.New("-partition is required when partitioning manually")
		}
	default:
		return errors.New(fmt.Sprintf("Partitioner %d not supported ", msg.Partition))
	}

	message := &sarama.ProducerMessage{Topic: msg.Topic, Partition: int32(msg.Partition)}

	if msg.Key != "" {
		message.Key = sarama.StringEncoder(msg.Key)
	}

	if msg.Value != "" {
		message.Value = sarama.StringEncoder(msg.Value)
	} else {
		return errors.New("messageValue is required")
	}

	if msg.UserName != "" {
		config.Net.SASL.Enable = true
		config.Net.SASL.User = msg.UserName
		config.Net.SASL.Password = msg.Password
	}

	producer, err := sarama.NewSyncProducer(strings.Split(msg.BrokerList, ","), config)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to open Kafka producer: %s", err))
	}
	defer func() {
		if err := producer.Close(); err != nil {
			fmt.Println("Failed to close Kafka producer cleanly:", err)
		}
	}()

	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to produce message: %s", err))
	} else { // if !*silent
		logs.Info("kafka:%v partition=%d\toffset=%d\n", msg, partition, offset)
	}
	return nil

}
