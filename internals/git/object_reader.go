package git

import (
	"bytes"
	"compress/zlib"
	"git-stats/internals/models"
	"io"
	"io/ioutil"
	"strings"
	"unicode"
)

// ReadObject ... Get contents of the object from a given path
func ReadObject(path string) (*models.CommitModel, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return readCompressedData(buf)
}

func removeSpace(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

func readCompressedData(buf []byte) (*models.CommitModel, error) {
	buffer := bytes.NewBuffer(buf)
	data, err := zlib.NewReader(buffer)

	if err != nil {
		return nil, err
	}

	return createGitCommitObject(data)
}

func extractString(data string, start string, end string) string {
	s := 0
	e := len(data)

	if start != "" {
		s = strings.Index(data, start)
		s += len(start)
	}

	if end != "" {
		e = strings.Index(data, end)
	}

	return strings.TrimSpace(data[s:e])
}

func extractTimeStamp(data string) string {
	splitData := strings.Split(data, " ")
	return splitData[len(splitData)-2]
}

func createGitCommitObject(data io.ReadCloser) (*models.CommitModel, error) {
	byteData, err := ioutil.ReadAll(data)
	if err != nil {
		return nil, err
	}

	data.Close()

	commitData := strings.Split(string(byteData), "\n")

	// for i := 0; i < len(commitData); i++ {
	// 	fmt.Println(i, " -> ", commitData[i])
	// }

	skipIndex := 0
	objectType := "commit"
	if strings.Index(commitData[2], "parent") >= 0 {
		skipIndex = 1
		objectType = "merge"
	}

	commitObj := &models.CommitModel{
		Commit: extractString(commitData[0], "commit", "tree"),
		Tree:   extractString(commitData[0], "tree", ""),
		Parent: extractString(commitData[1], "parent", ""),
		Author: &models.AuthorModel{
			Name:      extractString(commitData[skipIndex+2], "author", "<"),
			Email:     extractString(commitData[skipIndex+2], "<", ">"),
			TimeStamp: extractTimeStamp(commitData[skipIndex+2]),
		},
		Commiter: &models.CommiterModel{
			Name:      extractString(commitData[skipIndex+3], "committer", "<"),
			Email:     extractString(commitData[skipIndex+3], "<", ">"),
			TimeStamp: extractTimeStamp(commitData[skipIndex+3]),
		},
		Message: commitData[skipIndex+5],
		Type:    objectType,
	}

	return commitObj, nil
}
