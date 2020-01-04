package main

import "bygui86/go-zap-logger/logger"

func main() {

	logger.Logger.Debug("debug message")
	logger.Logger.Info("info message")
	logger.Logger.Warn("warn message")
	logger.Logger.Error("error message")
}
