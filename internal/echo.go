package internal

import "example/logger"

func Echo() {
	logger.NameInfof("is echo", "hi~~~ %d", 123456)
}
