package main

import (
	"bytes"
	"io"
	"net/http"
	"os/exec"
	"strings"
)

func reloadDnsmasq() error {
	err := createDnsmasqConfig()
	if err != nil {
		return err
	}
	return restartDnsmasq()
}

func restartDnsmasq() error {
	switch appConfig.DnsmasqReloaderConfig.Type {
	case "cmd":
		return restartDnsmasqByCmd(appConfig.DnsmasqReloaderConfig.cmdDnsmasqReloaderConfig)
	case "http":
		return restartDnsmasqByHTTP(appConfig.DnsmasqReloaderConfig.httpDnsmasqReloaderConfig)
	}
	return nil
}

func restartDnsmasqByCmd(opt cmdDnsmasqReloaderConfig) error {
	command := exec.Command("/bin/bash", "-c", opt.Command)
	var out bytes.Buffer
	command.Stderr = &out
	err := command.Run()
	if err != nil {
		return err
	}
	outErr := string(out.Bytes())
	if strings.Trim(outErr, " ") != "" {
		errorLogger.Printf("restart dnsmasq command %s stdout: %s", opt.Command, outErr)
	}
	return nil
}

func restartDnsmasqByHTTP(opt httpDnsmasqReloaderConfig) error {
	method := strings.ToUpper(opt.Method)
	if method == "" {
		method = "GET"
	}
	var reader io.Reader
	if opt.Body != "" {
		reader = strings.NewReader(opt.Body)
	}
	req, err := http.NewRequest(method, opt.URL, reader)
	if err != nil {
		return err
	}
	for k, v := range opt.Header {
		req.Header.Set(k, v)
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode >= 400 {
		errorLogger.Printf("restart dnsmasq http got response %d %s", res.StatusCode, res.Status)
	}
	return nil
}
