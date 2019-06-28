package test

import "github.com/zhao94254/common/logger"

func Fib(x int) int {
	if x < 2 {
		return x
	}
	logger.Log.Info(string(x))
	return Fib(x-1) + Fib(x-2)
}