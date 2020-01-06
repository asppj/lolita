package producer

import (
	"fmt"
	"sync"

	"t-mk-opentrace/pkg/mq/model"

	"t-mk-opentrace/config"

	kafka "git.dustess.com/mk-base/kafka-driver/producer"
	"git.dustess.com/mk-base/log"
)

// KafkaProducer 生产
type KafkaProducer struct {
	*kafka.Producer
}

var _normalProducer *KafkaProducer

var _normalOnce sync.Once

// DefaultProducer 生产
func DefaultProducer() *KafkaProducer {
	funcName := "producer-DefaultProducer"
	if _normalProducer != nil {
		return _normalProducer
	}
	_normalOnce.Do(func() {
		conf := config.Get().Kafka
		// TODO topic name
		topicName := conf.Topic.MKPlanReport.Name
		p := kafka.NewProducer(conf.Addrs, topicName)
		log.Info(fmt.Sprintf("%s-producer topic: %s", funcName, topicName))
		_normalProducer = &KafkaProducer{
			Producer: p,
		}
	})
	return _normalProducer
}

// AddOne 生产report
func AddOne(r model.Behave, companyID string, companyDB string) error {
	cr := struct {
		// TODO
	}{}
	_, _, err := DefaultProducer().ProduceOneJSON(cr)
	return err

}
