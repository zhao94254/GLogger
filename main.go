package main

import "github.com/zhao94254/common/logger"
import "github.com/zhao94254/common/test"

//var log = logger.InitLogger("./test.yml")

func initLog()  {
	logger.InitLogger("./test.yml")
}
func main()  {
	initLog()
	logger.Log.Info("test log")
	logger.Log.Debug("test log")
	logger.Log.Warn("test log")
	logger.Log.Error("test error")

	test.Fib(10)
}
