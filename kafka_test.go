package tools

import "testing"

func TestKafkaSender_KafkaCreateTopic(t *testing.T) {
	s := &KafkaSender{
		Uri:      "124.70.139.185:9092",
		UserName: "admin",
		Password: "admin001",
	}
	s.KafkaCreateTopic("test001")
}
