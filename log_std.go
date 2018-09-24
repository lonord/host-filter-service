// +build !syslog

package main

import (
	"log"
	"os"
)

var infoLogger *log.Logger
var warnLogger *log.Logger
var errorLogger *log.Logger

func init() {
	infoLogger = log.New(os.Stdout, "[INFO] ", log.LstdFlags)
	warnLogger = log.New(os.Stdout, "[WARN] ", log.LstdFlags)
	errorLogger = log.New(os.Stdout, "[ERROR] ", log.LstdFlags)
}
