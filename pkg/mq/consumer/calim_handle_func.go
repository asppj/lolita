package consumer

import (
	"git.dustess.com/mk-base/log"
	"github.com/go-errors/errors"
)

// DefaultHandlePanic 错误处理
func DefaultHandlePanic(payload interface{}) {
	funcName := "consumer-DefaultHandlePanic"
	if payload != nil {
		log.Err(funcName, errors.New(payload))
	} else {
		log.Err(funcName, errors.New("payload is nil"))
	}
}

// DefaultHandleFunc 消费逻辑
func DefaultHandleFunc(payload interface{}) {
	funcName := "consumer-DefaultHandleFunc"
	if param, ok := payload.(*DefaultHandlerParam); ok {
		c, err := param.ToModel()
		if err != nil {
			log.Err(funcName+"ToModel失败，须手动恢复", errors.New(err))
		}
		log.Info(c)

		param.Sess.MarkMessage(param.Payload, "done")
		log.Info("完成")
		return
	}
	log.Err(funcName, errors.New("Param格式错误。不会发生，发生就是写错了大bug"))
}
