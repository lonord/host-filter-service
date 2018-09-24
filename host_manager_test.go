package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteDnsmasqConfig(t *testing.T) {
	appSetting = &appSettingStruct{}
	err := resolveSetting([]byte(settingContent4DnsmasqConfig))
	assert.NoError(t, err)
	b := bytes.NewBuffer(make([]byte, 0))
	err = writeDnsmasqConfig(b)
	assert.NoError(t, err)
	assert.Equal(t, expectDnsmasqConfig, string(b.Bytes()))
}

const settingContent4DnsmasqConfig = `
{
	"standardHosts": [
		"aaa.com",
		"bbb.com",
		"ccc.ddd.abc"
	],
	"customHosts": [
		"qqq.com",
		"www.com",
		"bbb.com"
	],
	"customExcludeHosts": [
		"1.com",
		"2.com",
		"ccc.ddd.abc"
	],
	"standardHostsUpdateDate": 1537782713
}
`

const expectDnsmasqConfig = `address=/aaa.com/127.0.0.1
address=/bbb.com/127.0.0.1
address=/qqq.com/127.0.0.1
address=/www.com/127.0.0.1
`
