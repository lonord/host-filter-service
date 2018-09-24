package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func fetchStandardList() ([]string, error) {
	res, err := http.Get(defaultHostFetchURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return parseHostFormatContent(string(content)), nil
}

func parseHostFormatContent(content string) []string {
	lines := strings.Split(content, "\n")
	resultLines := []string{}
	for _, line := range lines {
		tl := strings.Trim(line, " ")
		if len(tl) == 0 {
			continue
		}
		if strings.HasPrefix(tl, "#") {
			continue
		}
		if strings.HasPrefix(tl, "127.0.0.1") {
			spaceIdx := strings.Index(tl, " ")
			if spaceIdx >= 0 && spaceIdx < len(tl) {
				resultLines = append(resultLines, substring(tl, spaceIdx+1))
			}
		}
	}
	return resultLines
}

func substring(source string, start int) string {
	var r = []rune(source)
	length := len(r)

	if start < 0 {
		return ""
	}

	if start == 0 {
		return source
	}

	var substring = ""
	for i := start; i < length; i++ {
		substring += string(r[i])
	}

	return substring
}
