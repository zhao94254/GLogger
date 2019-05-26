package main

import "common/logger"

var log = logger.InitLogger("./test.yml")

func main()  {
	log.Info("test log")
	log.Debug("test log")
	log.Warn("test log")
}
