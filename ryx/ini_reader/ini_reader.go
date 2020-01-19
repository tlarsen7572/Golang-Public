package ini_reader

import (
	"io/ioutil"
	"strings"
)

func LoadIni(path string) map[string]string {
	mapped := map[string]string{}
	ini, err := ioutil.ReadFile(path)

	if err != nil {
		return mapped
	}
	var stripped []byte
	for _, value := range ini {
		if value != 0 {
			stripped = append(stripped, value)
		}
	}

	settings := string(stripped)
	lines := strings.Split(settings, "\r\n")
	for _, line := range lines {
		split := strings.Split(line, `=`)
		if len(split) != 2 {
			continue
		}
		mapped[split[0]] = split[1]
	}
	return mapped
}
