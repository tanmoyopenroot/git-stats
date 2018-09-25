package git

import (
	"git-stats/internals/utils"
	"io/ioutil"
	"path"
	"strings"
)

func extractDataFromConfig(data []string, searchKey string) interface{} {
	for i := 0; i < len(data); i++ {
		if strings.Index(data[i], searchKey) >= 0 {
			switch searchKey {
			case "email":
				return utils.ExtractString(strings.TrimSpace(data[i]), "email = ", "")
			}
		}
	}

	switch searchKey {
	case "email":
		return ""
	default:
		return nil
	}
}

// GetUserEmail ... Get email of the user
func GetUserEmail(rootPath string) (string, error) {
	configPath := path.Join(rootPath, "config")

	buf, err := ioutil.ReadFile(configPath)
	if err != nil {
		return "", err
	}

	data := string(buf)
	configData := strings.Split(data, "\n")

	return extractDataFromConfig(configData, "email").(string), nil
}
