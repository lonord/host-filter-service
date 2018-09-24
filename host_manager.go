package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/emirpasic/gods/sets/linkedhashset"
)

func getCustomHosts() []string {
	return appSetting.getCustomHosts()
}

func addCustomHost(domain string) error {
	appSetting.customHosts.Add(domain)
	appSetting.customExcludeHosts.Remove(domain)
	err := saveSetting()
	if err != nil {
		return err
	}
	return nil
}

func removeCustomHost(domain string) error {
	appSetting.customHosts.Remove(domain)
	err := saveSetting()
	if err != nil {
		return err
	}
	return nil
}

func getCustomExcludeHosts() []string {
	return appSetting.getCustomExcludeHosts()
}

func addCustomExcludeHost(domain string) error {
	appSetting.customExcludeHosts.Add(domain)
	appSetting.customHosts.Remove(domain)
	err := saveSetting()
	if err != nil {
		return err
	}
	return nil
}

func removeCustomExcludeHost(domain string) error {
	appSetting.customExcludeHosts.Remove(domain)
	err := saveSetting()
	if err != nil {
		return err
	}
	return nil
}

func updateStandardHosts() error {
	list, err := fetchStandardList()
	if err != nil {
		return err
	}
	appSetting.standardHosts.Clear()
	for _, l := range list {
		appSetting.standardHosts.Add(l)
	}
	appSetting.standardHostsUpdateDate = uint64(time.Now().Unix())
	err = saveSetting()
	if err != nil {
		return err
	}
	return nil
}

func createDnsmasqConfig() error {
	filePath := path.Join(appConfig.DnsmasqConfigTargetDir, appConfig.DnsmasqConfigFileName)
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	err = writeDnsmasqConfig(f)
	if err != nil {
		return err
	}
	return nil
}

func deleteDnsmasqConfig() error {
	filePath := path.Join(appConfig.DnsmasqConfigTargetDir, appConfig.DnsmasqConfigFileName)
	err := os.Remove(filePath)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}

func writeDnsmasqConfig(writer io.Writer) error {
	hostSet := linkedhashset.New(appSetting.standardHosts.Values()...)
	hostSet.Add(appSetting.customHosts.Values()...)
	hostSet.Remove(appSetting.customExcludeHosts.Values()...)
	for _, h := range hostSet.Values() {
		_, err := fmt.Fprintf(writer, "address=/%s/127.0.0.1\n", h)
		if err != nil {
			return err
		}
	}
	return nil
}
