package git

import (
	"io/ioutil"
	"strings"
)

// GetSHA ... Get SHA from a given path
func GetSHA(path string) (string, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	data := strings.TrimSpace(string(buf))
	return data, nil
}
