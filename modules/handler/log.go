package handler

import (
	"log"
	"time"

	"github.com/naufalkhairil/Filewatcher/modules/event"
)

type logEvent struct{}

func NewLogHandler() *logEvent {
	return &logEvent{}
}

func (l *logEvent) HandleEvent(metadata event.EventMetadata) error {
	log.Printf("Event: File %s - Op %s - Size %d - Receive %s", metadata.Filename, metadata.Op, metadata.Size, metadata.TsReceive.Format(time.RFC3339))
	return nil
}
