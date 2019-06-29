package main

import (
	"github.com/zhao94254/common/logging"
)

//var log = logger.InitLogger("./test.yml")

//func initLog()  {
//	logger.InitLogger("./test.yml")
//}


func main()  {
	//initLog()
	//logger.Log.Info("test log")
	//logger.Log.Debug("test log")
	//logger.Log.Warn("test log")
	logging.InitLogger("logs/skr.log", "")
	logging.Infoln(logging.KV{"s":1}, "s")

	//test.Fib(10)

}
