package tools

import (
	"context"
	"errors"
	"github.com/astaxie/beego/logs"
	"parking_mall/databases"
	"parking_mall/databases/redis"
	"parking_mall/initial/config"
	"parking_mall/models/keyvalues"
	"parking_mall/utils/tools"
	"strconv"
	"testing"
	"time"
)

func TestNewQueue(t *testing.T) {
	databases.MysqlInit("mysql", "root:Helang@2022@tcp(47.244.160.194)/parking_mall?charset=utf8&parseTime=True&loc=Local")
	appKey := "parking_mall"
	rets := keyvalues.GetKeyValues(appKey)
	if rets == nil || len(rets) == 0 {
		panic("load config failed")
	}
	//databases.DbDefault.LogMode(CfgApp.DbLog)
	err := tools.DataToStruct(rets, &config.CfgGlobal)
	if err != nil {
		panic(err)
	}
	redis.InitRedis()

	config.CfgApp.IsMaster = true
	master := NewQueue()
	master.Register("test", MasterHandle)

	config.CfgApp.IsMaster = false
	slave := NewQueue()
	slave.Register("test", SlaveHandle)
	time.Sleep(time.Second)

	go func() {
		for i := 0; i < 1000; i += 2 {
			master.AddJob(&JobItem{
				Topic:   "test",
				Id:      strconv.Itoa(i),
				Seconds: int64(i),
				Message: strconv.Itoa(i),
			})
			logs.Info("slave push data ", i)
		}
	}()

	go func() {
		for i := 1; i < 1000; i += 2 {
			master.AddJob(&JobItem{
				Topic:   "test",
				Id:      strconv.Itoa(i),
				Seconds: int64(i),
				Message: strconv.Itoa(i),
			})
			logs.Info("master push data ", i)
		}
	}()

	time.Sleep(time.Hour)
}

func SlaveHandle(ctx context.Context, message string) error {
	select {
	case <-ctx.Done():
		return errors.New("超时")
	default:
		logs.Info("slave handle " + message)
		return nil
	}
}

func MasterHandle(ctx context.Context, message string) error {
	select {
	case <-ctx.Done():
		return errors.New("超时")
	default:
		logs.Info("master handle " + message)
		//return nil
		time.Sleep(time.Second * 2)
		return errors.New("测试")
	}
}
