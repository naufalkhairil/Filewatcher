package pubsub

import (
	"context"
	"sync"

	log "github.com/sirupsen/logrus"

	"cloud.google.com/go/pubsub"
	"gitlab.com/watonist/letsgo/bootstrap"
	"google.golang.org/api/option"
)

var client *pubsub.Client
var clientOnce sync.Once

var logger = log.WithField("modules", "pubsub")

func GetClient() *pubsub.Client {
	clientOnce.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), GetConnectTimeout())
		defer cancel()
		psClient, err := pubsub.NewClient(ctx, GetProject(), option.WithCredentialsFile(GetCredentialFile()))
		if err != nil {
			logger.WithError(err).Fatal("Unable to create pubsub client")
		}
		client = psClient

		bootstrap.AddShutdownRoutine(func() {
			if client == nil {
				return
			}

			logger.Debug("Closing pubsub client ...")
			if err := client.Close(); err != nil {
				logger.WithError(err).Fatal("Failed to close pubsub client")
			}
		})

	})

	return client
}
