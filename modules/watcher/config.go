package watcher

import (
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	cfgSourceDir       = "filewatcher.source-dir"
	cfgRefreshInterval = "filewatcher.refresh-interval"
	cfgHandlerType     = "filewatcher.handler"
	cfgEvents          = "filewatcher.events"
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

func GetWatchedEvents() fsnotify.Op {
	events := viper.GetStringSlice(cfgEvents)
	if len(events) == 0 {
		return fsnotify.Create | fsnotify.Write | fsnotify.Rename
	}

	var op fsnotify.Op
	for _, event := range events {
		switch event {
		case "create":
			op |= fsnotify.Create
		case "write":
			op |= fsnotify.Write
		case "rename":
			op |= fsnotify.Rename
		case "remove":
			op |= fsnotify.Remove
		case "chmod":
			op |= fsnotify.Chmod
		}
	}
	return op
}
