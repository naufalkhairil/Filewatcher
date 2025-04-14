package handler

import (
	"log"

	"github.com/naufalkhairil/Filewatcher/modules/event"
)

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
		log.Printf("Using %s handler", handlerType)
		return handler
	}

	log.Print("Using log handler")
	return HandlerMap["log"]
}
