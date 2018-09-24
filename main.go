package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {
	initConfig()
	initSetting()
	err := runWebServer()
	if err != nil {
		errorLogger.Fatalln("run web server error:", err)
	}
}

func runWebServer() error {
	err := reloadDnsmasq()
	if err != nil {
		return err
	}
	defer func() {
		deleteDnsmasqConfig()
		restartDnsmasq()
	}()
	w := newWebService()
	w.start()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan)
	for {
		select {
		case sig := <-signalChan:
			switch sig {
			case syscall.SIGHUP:
				fallthrough
			case syscall.SIGINT:
				fallthrough
			case syscall.SIGQUIT:
				fallthrough
			case syscall.SIGTERM:
				infoLogger.Printf("got singal %s, exit", sig.String())
				w.stop()
				return nil
			}
		}
	}
}
