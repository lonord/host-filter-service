package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type cmdDnsmasqReloaderConfig struct {
	Command string `yaml:"command"`
}

type httpDnsmasqReloaderConfig struct {
	URL    string            `yaml:"url"`
	Method string            `yaml:"method"`
	Header map[string]string `yaml:"header"`
	Body   string            `yaml:"body"`
}

type appConfigStruct struct {
	HostFetchURL           string `yaml:"hostFetchURL"`
	ServiceListenHost      string `yaml:"serviceListenHost"`
	ServiceListenPort      int    `yaml:"serviceListenPort"`
	DnsmasqConfigTargetDir string `yaml:"dnsmasqConfigTargetDir"`
	DnsmasqConfigFileName  string `yaml:"dnsmasqConfigFileName"`
	DnsmasqReloaderConfig  struct {
		Type                      string `yaml:"type"`
		cmdDnsmasqReloaderConfig  `yaml:",inline"`
		httpDnsmasqReloaderConfig `yaml:",inline"`
	} `yaml:"dnsmasqReloaderConfig"`
}

var appConfig appConfigStruct

func initConfig() {
	b, err := ioutil.ReadFile(standardConfigPath)
	if err != nil {
		warnLogger.Println("config file read error:", err, "using default")
	}
	err = resolveConfig(b)
	if err != nil {
		errorLogger.Fatalln("resolve config file error:", err)
	}
	err = checkConfig()
	if err != nil {
		errorLogger.Fatalln("config format error:", err)
	}
}

func resolveConfig(b []byte) error {
	err := yaml.Unmarshal(b, &appConfig)
	if err != nil {
		return err
	}
	if appConfig.HostFetchURL == "" {
		appConfig.HostFetchURL = defaultHostFetchURL
	}
	if appConfig.ServiceListenHost == "" {
		appConfig.ServiceListenHost = defaultListenHost
	}
	if appConfig.ServiceListenPort == 0 {
		appConfig.ServiceListenPort = defaultListenPort
	}
	if appConfig.DnsmasqConfigTargetDir == "" {
		appConfig.DnsmasqConfigTargetDir = defaultDnsmasqConfigTargetDir
	}
	if appConfig.DnsmasqConfigFileName == "" {
		appConfig.DnsmasqConfigFileName = defaultDnsmasqConfigFileName
	}
	return nil
}

func checkConfig() error {
	err := os.MkdirAll(appConfig.DnsmasqConfigTargetDir, os.ModePerm)
	if err != nil {
		return err
	}
	err = checkDnsmasqReloaderConfigType(appConfig.DnsmasqReloaderConfig.Type)
	if err != nil {
		return err
	}
	if appConfig.DnsmasqReloaderConfig.Type == "cmd" && appConfig.DnsmasqReloaderConfig.Command == "" {
		return fmt.Errorf("dnsmasqReloaderConfig.command is required when type is cmd")
	}
	if appConfig.DnsmasqReloaderConfig.Type == "http" && appConfig.DnsmasqReloaderConfig.URL == "" {
		return fmt.Errorf("dnsmasqReloaderConfig.url is required when type is http")
	}
	return nil
}

func checkDnsmasqReloaderConfigType(t string) error {
	if t != "cmd" && t != "http" {
		return fmt.Errorf("dnsmasqReloaderConfig.type should be cmd or http")
	}
	return nil
}
