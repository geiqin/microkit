package pubsub

import (
	"github.com/geiqin/gotools/helper"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/broker/nats"
	_ "github.com/micro/go-plugins/registry/nats/v2"
	"log"
	"os"
)

var myBroker broker.Broker

type EventType string

//消息注册
func init() {
	myBroker = nats.NewBroker(broker.Addrs(os.Getenv("MICRO_BROKER_ADDRESS")))
	myBroker.Init()
	err := myBroker.Connect()
	if err != nil {
		log.Println("broker connect error: %v", err)
	}
	log.Println("broker register succeed")
}

//消息发布
func Publish(eventName EventType, storeId int64, data string, headers ...map[string]string) error {
	heads := make(map[string]string)
	if storeId > 0 {
		heads["store_id"] = helper.Int64ToString(storeId)
	}
	if headers != nil {
		for k, v := range headers[0] {
			if k != "store_id" {
				heads[k] = v
			}
		}
	}
	msg := &broker.Message{
		Header: heads,
		Body:   []byte(data),
	}
	err := myBroker.Publish(string(eventName), msg)
	return err
}

//订阅消息
func Subscribe(eventName EventType, handler broker.Handler) (broker.Subscriber, error) {
	sub, err := myBroker.Subscribe(string(eventName), handler)
	return sub, err
}
