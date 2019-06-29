package test

import "github.com/zhao94254/common/logger"

func Fib(x int) int {
	if x < 2 {
		return x
	}
	logger.Log.Info(1)
	return Fib(x-1) + Fib(x-2)
}
