package main

import (
	"example/internal"
	"example/logger"
)

func main() {

	logger.Init()

	defer logger.Close()

	logger.Debugf("hello world")
	logger.Infof("hello world")
	logger.Errorf("hello world")

	internal.Hi()
	internal.Echo()

	internal.Run()
}
