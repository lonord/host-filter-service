package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/emirpasic/gods/sets/linkedhashset"
)

type settingStruct struct {
	StandardHosts           []string `json:"standardHosts"`
	CustomHosts             []string `json:"customHosts"`
	CustomExcludeHosts      []string `json:"customExcludeHosts"`
	StandardHostsUpdateDate uint64   `json:"standardHostsUpdateDate"`
}

type appSettingStruct struct {
	standardHosts           *linkedhashset.Set
	customHosts             *linkedhashset.Set
	customExcludeHosts      *linkedhashset.Set
	standardHostsUpdateDate uint64
}

var appSetting *appSettingStruct

func initSetting() {
	err := checkSettingDir()
	if err != nil {
		errorLogger.Fatalln("check setting directory error:", err)
	}
	err = loadSetting()
	if err != nil {
		errorLogger.Fatalln("load setting error:", err)
	}
}

func checkSettingDir() error {
	return os.MkdirAll(standardSettingDir, os.ModePerm)
}

func loadSetting() error {
	b, err := ioutil.ReadFile(standardSettingPath)
	if err != nil {
		return err
	}
	err = resolveSetting(b)
	if err != nil {
		return err
	}
	return nil
}

func resolveSetting(b []byte) error {
	s := settingStruct{}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	appSetting = &appSettingStruct{
		standardHosts:           newStringSet(s.StandardHosts),
		customHosts:             newStringSet(s.CustomHosts),
		customExcludeHosts:      newStringSet(s.CustomExcludeHosts),
		standardHostsUpdateDate: s.StandardHostsUpdateDate,
	}
	return nil
}

func (a *appSettingStruct) getStandardHosts() []string {
	return convertToStrings(a.standardHosts.Values())
}

func (a *appSettingStruct) getCustomHosts() []string {
	return convertToStrings(a.customHosts.Values())
}

func (a *appSettingStruct) getCustomExcludeHosts() []string {
	return convertToStrings(a.customExcludeHosts.Values())
}

func saveSetting() error {
	s := settingStruct{
		StandardHosts:           appSetting.getStandardHosts(),
		StandardHostsUpdateDate: appSetting.standardHostsUpdateDate,
		CustomHosts:             appSetting.getCustomHosts(),
		CustomExcludeHosts:      appSetting.getCustomExcludeHosts(),
	}
	b, err := json.Marshal(s)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(standardSettingPath, b, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func convertToStrings(params []interface{}) []string {
	strArray := make([]string, len(params))
	for i, arg := range params {
		strArray[i] = arg.(string)
	}
	return strArray
}

func newStringSet(strs []string) *linkedhashset.Set {
	s := linkedhashset.New()
	for _, str := range strs {
		s.Add(str)
	}
	return s
}
