package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	appName                       = "__unknow__"
	appVersion                    = ""
	buildTime                     = ""
	configPathPrefix              = "/usr/local"
	defaultHostFetchURL           = "https://raw.githubusercontent.com/vokins/yhosts/master/hosts"
	standardConfigPath            = configPathPrefix + "/" + appName + "/config.yml"
	standardSettingDir            = "/var/local/" + appName
	standardSettingPath           = standardSettingDir + "/setting.json"
	defaultListenHost             = "0.0.0.0"
	defaultListenPort             = 2019
	defaultDnsmasqConfigTargetDir = "/etc/dnsmasq.d"
	defaultDnsmasqConfigFileName  = "host_filter.conf"
)

func init() {
	checkExecValid()
	var showVer bool
	flag.BoolVar(&showVer, "v", false, "show build version")
	flag.Parse()
	if showVer {
		fmt.Printf("build ver:\t%s\n", appVersion)
		fmt.Printf("build time:\t%s\n", buildTime)
		os.Exit(0)
	}
}

func checkExecValid() {
	if appName == "__unknow__" {
		log.Fatalln("invalid exectuable: build flags error")
	}
}
