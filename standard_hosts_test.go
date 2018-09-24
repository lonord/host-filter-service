package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubstring(t *testing.T) {
	s := "hello world"
	rs := substring(s, 6)
	assert.Equal(t, "world", rs)
}

func TestSubstring2(t *testing.T) {
	s := "你好，世界"
	rs := substring(s, 3)
	assert.Equal(t, "世界", rs)
}

func TestParseHostFormatContent(t *testing.T) {
	strs := parseHostFormatContent(hostContent)
	assert.Equal(t, []string{
		"activate.adobe.com",
		"ereg.adobe.com",
		"hlrcv.stage.adobe.com",
		"lm.licenses.adobe.com",
		"lmlicenses.wip4.adobe.com",
		"na1r.services.adobe.com",
		"na2m-pr.licenses.adobe.com",
		"serial.alcohol-soft.com",
		"trial.alcohol-soft.com",
	}, strs)
}

const hostContent = `
#version=201809241807
#url=https://github.com/vokins/yhosts
127.0.0.1 activate.adobe.com
127.0.0.1 ereg.adobe.com
127.0.0.1 hlrcv.stage.adobe.com
127.0.0.1 lm.licenses.adobe.com
127.0.0.1 lmlicenses.wip4.adobe.com
127.0.0.1 na1r.services.adobe.com
127.0.0.1 na2m-pr.licenses.adobe.com
127.0.0.1 serial.alcohol-soft.com
127.0.0.1 trial.alcohol-soft.com
`
