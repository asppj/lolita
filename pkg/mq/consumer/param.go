package consumer

import (
	"encoding/json"

	"t-mk-opentrace/pkg/mq/model"

	"github.com/Shopify/sarama"
	"github.com/go-errors/errors"
)

// DefaultHandlerParam 参数
type DefaultHandlerParam struct {
	Payload *sarama.ConsumerMessage
	Sess    sarama.ConsumerGroupSession
	// Type    string
}

// ToModel 转换
func (p *DefaultHandlerParam) ToModel() (*model.Behave, error) {
	if p.Payload == nil {
		return nil, errors.New("consumer-DefaultHandlerParam-Payload is nil")
	}
	// TODO
	c := &model.Behave{}
	if err := json.Unmarshal(p.Payload.Value, c); err != nil {
		return nil, err
	}
	return c, nil
}
