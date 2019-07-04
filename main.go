package main

import (
	"fmt"
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

	logging.InitLogger("logs/skr.log", logging.Conf{"rotation":"hour"})
	for i:=0;i<5 ;i++  {
		logging.Infoln(logging.KV{"dp": 1, "dpp": 1.0, "dppp": 5, "t": 100}, "s")
	}
	logging.Errorln(nil, logging.KV{"dp": 1, "dpp": 1.0, "dppp": 5, "t": 100}, "a")
	fmt.Println("a", logging.KV{"dp": 1, "dpp": 1.0, "dppp": 5, "t": 100})
	//test.Fib(10)

}
