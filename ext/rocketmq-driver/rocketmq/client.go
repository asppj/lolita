package rocketmq

import (
	"context"
	"fmt"
	
	rocketmq "github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

const (
	topic = "testproducer"
	addr  = "127.0.0.1:10911"
)

var (
	_producer     rocketmq.Producer
	_pushConsumer rocketmq.PushConsumer
)

func InitProducer() (err error) {
	endPoint, err := primitive.NewNamesrvAddr(addr)
	if err != nil {
		return err
	}
	_producer, err := rocketmq.NewProducer(
		producer.WithNameServer(endPoint),
		// producer.WithNsResolver(primitive.NewPassthroughResolver(endPoint)),
		producer.WithRetry(2),
		producer.WithGroupName("GID_xxxxxx"),
	)
	if err != nil {
		return err
	}
	err = _producer.Start()
	if err != nil {
		_producer = nil
	}
	return
}

func SendOneWay(ctx context.Context, msg string) error {
	return _producer.SendOneWay(ctx, &primitive.Message{
		Topic: "topic",
		Body:  []byte(msg),
		// Flag:          0,
		// TransactionId: "",
		// Batch:         false,
		// Queue:         &primitive.MessageQueue{},
	})
}

func InitPushConsumer() (err error) {
	endPoint, err := primitive.NewNamesrvAddr(addr)
	if err != nil {
		return err
	}
	_pushConsumer, err = rocketmq.NewPushConsumer(
		consumer.WithNameServer(endPoint),
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithGroupName("GID_XXXXXX"),
	)
	if err != nil {
		return
	}
	// Subscribe
	_pushConsumer.Subscribe(topic, consumer.MessageSelector{}, func(c context.Context, me ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		fmt.Printf("sub:%v", me)
		return consumer.ConsumeSuccess, nil
	})
	
	// start
	err = _pushConsumer.Start()
	if err != nil {
		_pushConsumer = nil
	}
	return err
}

type ConsumerEndpointer interface {
	Consumer(ctx context.Context, payload []byte) error
}
