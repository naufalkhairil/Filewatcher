# Filewatcher

A Go-based file monitoring tool that watches directory changes and publishes events to various handlers including Google Cloud Pub/Sub and logging.

## Features

- **Real-time file monitoring** - Watches for file create, write, and rename operations
- **Multiple handlers** - Support for logging and Google Cloud Pub/Sub event publishing
- **Configurable** - YAML-based configuration with environment variable support
- **Event validation** - Built-in file hash validation to prevent duplicate processing
- **Debounced events** - Intelligent event batching to handle rapid file changes

## Installation

```bash
go install github.com/naufalkhairil/Filewatcher@latest
```

Or build from source:

```bash
git clone https://github.com/naufalkhairil/Filewatcher.git
cd Filewatcher
go build -o filewatcher
```

## Configuration

Create a `filewatcher.yaml` configuration file:

```yaml
pubsub:
  credential: credential.json
  project: your-gcp-project
  topic: your-pubsub-topic

filewatcher:
  source-dir: /path/to/watch
  handler: pubsub # supported: log, pubsub

log:
  level: info # debug, info, warn, error
```

### Configuration Options

- `pubsub.credential`: Path to GCP service account JSON file
- `pubsub.project`: Google Cloud Project ID
- `pubsub.topic`: Pub/Sub topic name
- `filewatcher.source-dir`: Directory to monitor
- `filewatcher.handler`: Event handler type (`log` or `pubsub`)
- `log.level`: Logging level

## Usage

Start the file watcher:

```bash
filewatcher watcher
```

The tool will:
1. Monitor the configured source directory
2. Detect file changes (create, write, rename)
3. Generate event metadata with file hash validation
4. Send events to the configured handler

## Event Structure

Events published to handlers contain:

```json
{
  "file_path": "/path/to/changed/file",
  "operation": "CREATE|WRITE|RENAME",
  "timestamp": "2025-02-05T14:41:19Z",
  "hash": "sha256-hash-of-file",
  "size": 1024
}
```

## Handlers

### Log Handler
Outputs events to stdout/stderr with structured logging.

### Pub/Sub Handler
Publishes events as JSON messages to Google Cloud Pub/Sub topics.

## Architecture

```
├── cmd/                 # CLI commands
├── modules/
│   ├── client/pubsub/   # Google Cloud Pub/Sub client
│   ├── event/           # Event metadata generation
│   ├── handler/         # Event handlers (log, pubsub)
│   ├── validator/       # File validation and hashing
│   └── watcher/         # File system monitoring
└── main.go             # Application entry point
```

## Development

Requirements:
- Go 1.23.2+
- Google Cloud SDK (for Pub/Sub handler)

Run tests:
```bash
go test ./...
```

## License

Licensed under the terms specified in the LICENSE file.

## Author

Naufal Khairil Imami