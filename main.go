package main

import (
	"time"

	"github.com/Sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
	var count int
	for {
		count++
		logger.Info(count)
		time.Sleep(1 * time.Second)
	}
}
