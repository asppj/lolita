package consumer

import (
	"fmt"
	"sync"
	"time"

	"t-mk-opentrace/pkg/mq/model"

	"t-mk-opentrace/config"

	kafka "git.dustess.com/mk-base/kafka-driver/consumer"
	"git.dustess.com/mk-base/log"
	"github.com/Shopify/sarama"
	"github.com/panjf2000/ants/v2"
)

const (
	beType = model.BeType
)

// KafkaConsumer 微信事件消费者
type KafkaConsumer struct {
	*kafka.Consumer
}

var _defaultConsumer *KafkaConsumer
var _defaultConsumerLock sync.Mutex

// DefaultConsumer 默认行为事件消费者
func DefaultConsumer() *KafkaConsumer {
	if _defaultConsumer != nil {
		return _defaultConsumer
	}

	_defaultConsumerLock.Lock()
	if _defaultConsumer == nil {
		kConf := config.Get().Kafka

		// config.Get().Kafka.Group.MKPmFDBLoad
		// kafka config addr Group Topic
		// TODO topicname
		c := kafka.NewConsumer(kConf.Addrs, kConf.Topic.MKPlanReport.Name, kConf.Group.MKWatWXLoad, &DefaultHandler{
			poolSize: 10,
		})
		_defaultConsumer = &KafkaConsumer{Consumer: c}
		log.Info(fmt.Sprintf("consumer topic： %s, consumer group: %s", c.Topic, c.Group))
	}
	_defaultConsumerLock.Unlock()

	return _defaultConsumer
}

// DefaultHandler 微信事件处理器
type DefaultHandler struct {
	pool     *ants.PoolWithFunc
	poolSize int
}

// Setup is run at the beginning of a new session, before ConsumeClaim.
func (beh *DefaultHandler) Setup(s sarama.ConsumerGroupSession) error {
	log.Info("consumer setup")
	var err error
	beh.pool, err = ants.NewPoolWithFunc(beh.poolSize, DefaultHandleFunc,
		ants.WithPanicHandler(DefaultHandlePanic))
	return err
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
// but before the offsets are committed for the very last time.
func (beh *DefaultHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	log.Info("consumer cleanup")
	for {
		running := beh.pool.Running()
		if running > 0 {
			log.Warn("handler thread is running, check after 1s", running)
			time.Sleep(time.Second * 1)
		}
		break
	}
	defer beh.pool.Release()
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// Once the Messages() channel is closed, the Handler must finish its processing
// loop and exit.
func (beh *DefaultHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// 协程池
	for payload := range claim.Messages() {
		err := beh.pool.Invoke(&DefaultHandlerParam{
			Payload: payload,
			Sess:    sess,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
