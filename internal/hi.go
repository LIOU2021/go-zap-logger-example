package internal

import "example/logger"

func Hi() {
	logger.Infof("hi %s", "hello")
}
