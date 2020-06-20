package tools

import (
	"context"
	"errors"
	"github.com/astaxie/beego/logs"
	"parking_mall/databases/redis"
	"parking_mall/initial/config"
	"sync"
	"time"
)

var GlobalDelayQueue *DelayQueue

type Callback func(ctx context.Context, message string) error

type DelayQueue struct {
	Handles map[string]Callback
}

func NewQueue() *DelayQueue {
	q := new(DelayQueue)
	q.Handles = make(map[string]Callback)

	if config.CfgApp.IsMaster {
		go q.DelayHandle()
	}
	go q.ReadyHandle()
	return q
}

func (i *DelayQueue) Register(topic string, call Callback) {
	i.Handles[topic] = call
}

type JobItem struct {
	Topic   string
	Id      string
	Seconds int64
	Message string
}

func (i *DelayQueue) AddJob(job *JobItem) error {
	if job == nil {
		return errors.New("job is nil")
	}
	if i.Handles[job.Topic] == nil {
		return errors.New("topic:" + job.Topic + " 未注册方法")
	}
	if r, _ := redis.HExistsBool(job.Topic, job.Id); r {
		return errors.New("job: " + job.Id + " existed")
	}

	delayTime := time.Now().Unix() + job.Seconds

	err := redis.ZAdd(job.Topic, delayTime, job.Id)
	if err != nil {
		return err
	}
	err = redis.HSetInterface(job.Topic+"hash", job.Id, job.Message)
	if err != nil {
		redis.ZDelete(job.Topic, job.Id)
		return err
	}
	return nil
}

// 将每个topic内的超时job都放入待处理list
func (i *DelayQueue) DelayHandle() {
	wg := new(sync.WaitGroup)
	for k := range i.Handles {
		wg.Add(1)
		go func(topic string) {
			defer wg.Done()
			handle := i.Handles[topic]
			if handle == nil {
				return
			}
			t := time.NewTicker(time.Second * 2)
			defer t.Stop()
			for {
				<-t.C
				now := time.Now().Unix()
				vals, _ := redis.ZRangeByScore(topic, 0, now)
				if len(vals) != 0 {
					redis.ZRemRangeByScore(topic, 0, now)
					for id := range vals {
						logs.Info("redis list push ", id)
						redis.LPush(topic+"list", id)
					}
				}
			}
		}(k)
	}
	wg.Wait()
}

func (i *DelayQueue) ReadyHandle() {
	wg := new(sync.WaitGroup)
	for k := range i.Handles {
		wg.Add(1)
		go func(topic string) {
			defer wg.Done()
			handle := i.Handles[topic]
			if handle == nil {
				return
			}
			t := time.NewTicker(time.Millisecond * 100)
			defer t.Stop()
			for {
				<-t.C
				jobId, _ := redis.LPop(topic + "list")
				if jobId != "" {
					message, _ := redis.HGetString(topic+"hash", jobId)
					ctx, cancle := context.WithTimeout(context.Background(), time.Second)
					err := handle(ctx, message)
					if err != nil {
						// 处理失败 需要归还key
						logs.Info("添加延时: ", jobId)
						err := redis.ZAdd(topic, time.Now().Unix()+5, jobId)
						if err != nil {
							logs.Error("延时队列 zadd topic:%s failed", topic)
						}
					} else {
						// 成功 删除
						redis.HDel(topic+"hash", jobId)
					}
					cancle()
				}
			}
		}(k)
	}
	wg.Wait()
}
