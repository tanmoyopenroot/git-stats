package git

import (
	"io/ioutil"
	"path"
	"strings"
)

func processPath(path string) string {
	if len(path) == 0 {
		return ""
	}

	headPath := strings.SplitAfter(path, "ref: ")[1]
	return strings.TrimSpace(headPath)
}

// CurrentBranchPath ... Get current branch path of the repo
func CurrentBranchPath(rootPath string) (string, error) {
	headPath := path.Join(rootPath, "HEAD")

	buf, err := ioutil.ReadFile(headPath)
	if err != nil {
		return "", err
	}

	data := string(buf)
	return processPath(data), nil
}
