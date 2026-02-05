package watcher

import (
	"os"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/fsnotify/fsnotify"
	eventModules "github.com/naufalkhairil/Filewatcher/modules/event"
	handlerModules "github.com/naufalkhairil/Filewatcher/modules/handler"
)

var watcher *fsnotify.Watcher
var watcherDone chan bool
var watcherWaitGroup sync.WaitGroup
var handler handlerModules.Handler

var pendingEvents map[string]*time.Timer

var logger = log.WithField("modules", "watcher")

func Start() error {
	logger.WithField("source", GetSourceDir()).Info("Starting filewatcher ...")

	watcherDone = make(chan bool)
	pendingEvents = make(map[string]*time.Timer)
	handler = handlerModules.GetHandler(GetHandlerType())

	newWatcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	watcher = newWatcher
	defer watcher.Close()

	go ProcessEvents()

	watcherWaitGroup.Add(1)
	go func(watcher *fsnotify.Watcher) {
		defer watcherWaitGroup.Done()
		for {
			if err := watcher.Add(GetSourceDir()); err != nil {
				log.WithError(err).WithField("source", GetSourceDir()).Fatal("Failed to watch directory")
			}
		}
	}(watcher)

	watcherWaitGroup.Wait()
	return nil
}

func ProcessEvents() {
	watcherWaitGroup.Add(1)
	defer watcherWaitGroup.Done()

	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	for {

		select {
		case event, ok := <-watcher.Events:
			if !ok {
				log.Printf("Failed to process event, %s op: %s", event.Name, event.Op.String())
				continue
			}

			log.Printf("Received event: %s op: %s", event.Name, event.Op.String())

			if event.Op&(fsnotify.Create|fsnotify.Write|fsnotify.Rename) != 0 {
				handleEvent(event)
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Fatalf("Watcher error, %s", err)
		case _, ok := <-ticker.C:
			if !ok {
				log.Print("Ticker channel closed")
				continue
			}
		}
	}
}

func handleEvent(event fsnotify.Event) {

	if timer, exists := pendingEvents[event.Name]; exists {
		timer.Stop()
	}

	pendingEvents[event.Name] = time.AfterFunc(GetRefreshInterval(), func() {
		if _, err := os.Stat(event.Name); os.IsNotExist(err) {
			if event.Op == fsnotify.Rename {
				// log.Printf("File %s is move or deleted", event.Name)
				delete(pendingEvents, event.Name)
				return
			}
		}

		// valid, err := validator.ValidateFile(event.Name)
		// if err != nil {
		// 	log.Printf("Failed to validate file %s, %s", event.Name, err)
		// 	return
		// }

		// if valid {
		// 	log.Printf("File %s is valid", event.Name)
		// }

		eventMeta, err := eventModules.GenerateMetadata(event)
		if err != nil {
			log.Printf("Failed to generate metadata for file %s, %s", event.Name, err)
			delete(pendingEvents, event.Name)
			return
		}

		handler.HandleEvent(eventMeta)
		// log.Printf("Event metadata: %v", eventMeta)

		delete(pendingEvents, event.Name)
	})
}
