package event

import (
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
)

type EventMetadata struct {
	Filename  string
	Op        string
	Size      int
	TsReceive time.Time
}

func GenerateMetadata(event fsnotify.Event) (EventMetadata, error) {
	metadata := EventMetadata{
		Filename:  event.Name,
		Op:        event.Op.String(),
		TsReceive: time.Now(),
	}

	fileInfo, err := os.Stat(event.Name)
	if err != nil {
		return metadata, err
	}
	metadata.Size = int(fileInfo.Size())

	return metadata, nil
}
