// +build syslog

package main

import (
	"log"
	"log/syslog"
)

var infoLogger *log.Logger
var warnLogger *log.Logger
var errorLogger *log.Logger

func init() {
	logger, err := syslog.NewLogger(syslog.LOG_INFO, log.LstdFlags)
	if err != nil {
		log.Fatal(err)
	}
	infoLogger = logger

	logger, err = syslog.NewLogger(syslog.LOG_WARNING, log.LstdFlags)
	if err != nil {
		log.Fatal(err)
	}
	warnLogger = logger

	logger, err = syslog.NewLogger(syslog.LOG_ERR, log.LstdFlags)
	if err != nil {
		log.Fatal(err)
	}
	errorLogger = logger
}
