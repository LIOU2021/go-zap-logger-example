package internal

import (
	"example/logger"
	"time"

	"github.com/google/uuid"
)

func Run() {

	for {
		id := uuid.New().String()
		logger.NameInfof(id, "test | x-api-id: %s", id)
		time.Sleep(1 * time.Second)
	}
}
