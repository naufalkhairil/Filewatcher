package handler

import (
	"context"
	"encoding/json"
	"log"

	"cloud.google.com/go/pubsub"
	pubsubClient "github.com/naufalkhairil/Filewatcher/modules/client/pubsub"
	"github.com/naufalkhairil/Filewatcher/modules/event"
)

type pubsubEvent struct{}

func NewPubsubHandler() *pubsubEvent {
	return &pubsubEvent{}
}

func (p *pubsubEvent) HandleEvent(metadata event.EventMetadata) error {
	client := pubsubClient.GetClient()
	topic := client.Topic(pubsubClient.GetTopic())

	data, err := json.Marshal(metadata)
	if err != nil {
		return err
	}

	ctx := context.Background()
	res := topic.Publish(ctx, &pubsub.Message{
		Data: data,
	})

	if _, err := res.Get(ctx); err != nil {
		return err
	}

	log.Printf("Published event to topic %s", topic)
	return nil
}
