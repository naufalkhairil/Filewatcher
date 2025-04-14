package pubsub

import (
	"context"
	"log"
	"sync"

	"cloud.google.com/go/pubsub"
	"gitlab.com/watonist/letsgo/bootstrap"
	"google.golang.org/api/option"
)

var client *pubsub.Client
var clientOnce sync.Once

func GetClient() *pubsub.Client {
	clientOnce.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), GetConnectTimeout())
		defer cancel()
		psClient, err := pubsub.NewClient(ctx, GetProject(), option.WithCredentialsFile(GetCredentialFile()))
		if err != nil {
			log.Fatalf("Unable to create pubsub client, %s", err)
		}
		client = psClient

		bootstrap.AddShutdownRoutine(func() {
			if client == nil {
				return
			}

			log.Println("Closing pubsub client ...")
			if err := client.Close(); err != nil {
				log.Fatalf("Failed to close pubsub client, %s", err)
			}
		})

	})

	return client
}
