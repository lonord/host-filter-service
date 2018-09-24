package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResolveConfig(t *testing.T) {
	appConfig = appConfigStruct{}
	err := resolveConfig([]byte(configContent))
	assert.NoError(t, err)
	assert.Equal(t, "http://www.aaa.com", appConfig.HostFetchURL)
	assert.Equal(t, "127.0.0.1", appConfig.ServiceListenHost)
	assert.Equal(t, 12345, appConfig.ServiceListenPort)
	assert.Equal(t, "/var/abc", appConfig.DnsmasqConfigTargetDir)
	assert.Equal(t, "aaa.conf", appConfig.DnsmasqConfigFileName)
	assert.Equal(t, "cmd", appConfig.DnsmasqReloaderConfig.Type)
	assert.Equal(t, "ls -alh", appConfig.DnsmasqReloaderConfig.Command)
}

func TestResolveConfig2(t *testing.T) {
	appConfig = appConfigStruct{}
	err := resolveConfig([]byte(configContent2))
	assert.NoError(t, err)
	assert.Equal(t, "http", appConfig.DnsmasqReloaderConfig.Type)
	assert.Equal(t, "http://localhost:11002/dnsmasq", appConfig.DnsmasqReloaderConfig.URL)
	assert.Equal(t, "PUT", appConfig.DnsmasqReloaderConfig.Method)
	assert.Equal(t, "", appConfig.DnsmasqReloaderConfig.Body)
	assert.Nil(t, appConfig.DnsmasqReloaderConfig.Header)
}

func TestResolveConfig3(t *testing.T) {
	appConfig = appConfigStruct{}
	err := resolveConfig([]byte(configContent3))
	assert.NoError(t, err)
	assert.Equal(t, "https://raw.githubusercontent.com/vokins/yhosts/master/hosts", appConfig.HostFetchURL)
	assert.Equal(t, "0.0.0.0", appConfig.ServiceListenHost)
	assert.Equal(t, 2019, appConfig.ServiceListenPort)
	assert.Equal(t, "/etc/dnsmasq.d", appConfig.DnsmasqConfigTargetDir)
	assert.Equal(t, "host_filter.conf", appConfig.DnsmasqConfigFileName)
	assert.Equal(t, "cmd", appConfig.DnsmasqReloaderConfig.Type)
	assert.Equal(t, "service dnsmasq restart", appConfig.DnsmasqReloaderConfig.Command)
}

const configContent = `
hostFetchURL: http://www.aaa.com
serviceListenHost: 127.0.0.1
serviceListenPort: 12345
dnsmasqConfigTargetDir: /var/abc
dnsmasqConfigFileName: aaa.conf
dnsmasqReloaderConfig:
  type: cmd
  command: ls -alh
`

const configContent2 = `
hostFetchURL: http://www.aaa.com
serviceListenHost: 127.0.0.1
serviceListenPort: 12345
dnsmasqConfigTargetDir: /var/abc
dnsmasqConfigFileName: aaa.conf
dnsmasqReloaderConfig:
  type: http
  url: http://localhost:11002/dnsmasq
  method: PUT
`

const configContent3 = `
###############################################################################
#                                                                             #
#                  configure file of host filter service                      #
#                                                                             #
###############################################################################

# fetch url of hosts formatted list file, default https://raw.githubusercontent.com/vokins/yhosts/master/hosts
#hostFetchURL: https://raw.githubusercontent.com/vokins/yhosts/master/hosts

# service server listen host, default 0.0.0.0
#serviceListenHost: '0.0.0.0'

# service server listen port, default 2019
#serviceListenPort: 2019

# directory of dnsmasq config file to generate to, default /etc/dnsmasq.d
#dnsmasqConfigTargetDir: /etc/dnsmasq.d

# name of dnsmasq config file to generate, default host_filter.conf
#dnsmasqConfigFileName: host_filter.conf

# config of dnsmasq reloading
# <1> reloading by shell command (default)
dnsmasqReloaderConfig:
  type: 'cmd'
  command: 'service dnsmasq restart'
# <2> reloading by http request
#dnsmasqReloaderConfig:
#  type: 'http'
#  url: 'http://yourdomain/path'
#  method: 'POST'
#  header:
#    headerkey: 'headervalue'
#  body: 'data'
`
