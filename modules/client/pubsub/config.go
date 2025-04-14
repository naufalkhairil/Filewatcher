package pubsub

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

const (
	credential     = "pubsub.credential"
	project        = "pubsub.project"
	connectTimeout = "pubsub.connect-timeout"
	topic          = "pubsub.topic"
)

func GetCredentialFile() string {
	cred := viper.GetString(credential)
	if cred == "" {
		log.Fatalf("Failed to get credential, %s", credential)
	}

	return cred
}

func GetProject() string {
	return viper.GetString(project)
}

func GetTopic() string {
	topic := viper.GetString(topic)
	if topic == "" {
		log.Fatalf("Failed to get topic, %s", topic)
	}

	return topic
}

func GetConnectTimeout() time.Duration {
	timeout := viper.GetDuration(connectTimeout)
	if timeout == 0 {
		return 10 * time.Second
	}

	return timeout
}
