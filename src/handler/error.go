package handler

import "merger/src/logger"

func FailOnError(err error, message string) {
	if err != nil {
		logger.Error(err.Error())
		logger.Info(message)
		panic(err)
	}
}

func Fail(message string) {
	logger.Error(message)
	panic(message)
}
