package handler

import (
	log "github.com/sirupsen/logrus"

	"github.com/naufalkhairil/Filewatcher/modules/event"
)

var logger = log.WithField("modules", "handler")

var HandlerMap map[string]Handler = make(map[string]Handler)

type Handler interface {
	HandleEvent(metadata event.EventMetadata) error
}

func InitHandler() {
	HandlerMap["log"] = NewLogHandler()
	HandlerMap["pubsub"] = NewPubsubHandler()
}

func GetHandler(handlerType string) Handler {
	if handler, ok := HandlerMap[handlerType]; ok {
		logger.Infof("Using %s handler", handlerType)
		return handler
	}

	logger.Info("Using log handler")
	return HandlerMap["log"]
}
