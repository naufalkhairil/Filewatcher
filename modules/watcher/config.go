package watcher

import (
	"time"

	"github.com/spf13/viper"
)

const (
	cfgSourceDir       = "filewatcher.source-dir"
	cfgRefreshInterval = "filewatcher.refresh-interval"
	cfgHandlerType     = "filewatcher.handler"
)

func GetSourceDir() string {
	return viper.GetString(cfgSourceDir)
}

func GetRefreshInterval() time.Duration {
	duration := viper.GetDuration(cfgRefreshInterval)
	if duration == time.Duration(0) {
		return 5 * time.Second
	}

	return duration
}

func GetHandlerType() string {
	handlerType := viper.GetString(cfgHandlerType)
	if handlerType == "" {
		return "log"
	}

	return handlerType
}
