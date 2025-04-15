package validator

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

var logger = log.WithField("modules", "validator")

func ValidateFile(filePath string) (bool, error) {
	logger.WithField("file", filePath).Debug("Validating file ...")

	fileHash, err := generateHash(filePath)
	if err != nil {
		return false, err
	}

	retries := 10
	attempt := 0
	for {
		if attempt >= retries {
			return false, fmt.Errorf("File %s is not valid after %d attempts", filePath, retries)
		}

		// Refreshing hash
		time.Sleep(1 * time.Second)
		logger.Debug("Refreshing hash ...")
		freshHash, err := generateHash(filePath)
		if err != nil {
			return false, err
		}

		if fileHash == freshHash {
			return true, nil
		}

		fileHash = freshHash

		attempt++
	}
}

func generateHash(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
