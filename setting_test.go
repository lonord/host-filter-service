package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResolveSetting(t *testing.T) {
	appSetting = &appSettingStruct{}
	err := resolveSetting([]byte(settingContent))
	assert.NoError(t, err)
	assert.Equal(t, []string{"aaa.com", "bbb.com", "ccc.ddd.abc"}, appSetting.getStandardHosts())
	assert.Equal(t, []string{"qqq.com", "www.com", "eee.com"}, appSetting.getCustomHosts())
	assert.Equal(t, []string{"1.com", "2.com", "3.com"}, appSetting.getCustomExcludeHosts())
	assert.Equal(t, uint64(1537782713), appSetting.standardHostsUpdateDate)
}

const settingContent = `
{
	"standardHosts": [
		"aaa.com",
		"bbb.com",
		"ccc.ddd.abc"
	],
	"customHosts": [
		"qqq.com",
		"www.com",
		"eee.com"
	],
	"customExcludeHosts": [
		"1.com",
		"2.com",
		"3.com"
	],
	"standardHostsUpdateDate": 1537782713
}
`
