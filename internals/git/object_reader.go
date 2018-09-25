package git

import (
	"bytes"
	"compress/zlib"
	"git-stats/internals/models"
	"io"
	"io/ioutil"
	"strings"
)

// ReadObject ... Get contents of the object from a given path
func ReadObject(path string) (*models.CommitModel, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return readCompressedData(buf)
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

func extractCommiter(data string) models.Commiter {
	return models.Commiter{
		Name:      extractString(data, "committer", "<"),
		Email:     extractString(data, "<", ">"),
		TimeStamp: extractTimeStamp(data),
	}
}

func extractAuthor(data string) models.Author {
	return models.Author{
		Name:      extractString(data, "author", "<"),
		Email:     extractString(data, "<", ">"),
		TimeStamp: extractTimeStamp(data),
	}
}

func isTypeMerge(data []string) bool {
	return strings.Index(data[2], "parent") > 0
}

func extractDataFromCommit(data []string, searchKey string) interface{} {
	for i := 0; i < len(data); i++ {
		if strings.Index(data[i], searchKey) >= 0 {
			switch searchKey {
			case "tree":
				return extractString(data[i], "tree", "")
			case "parent":
				return extractString(data[i], "parent", "")
			case "author":
				return extractAuthor(data[i])
			case "commiter":
				return extractCommiter(data[i])
			}
		}
	}

	switch searchKey {
	case "tree":
		return ""
	case "parent":
		return ""
	case "author":
		return models.Author{}
	case "commiter":
		return models.Commiter{}
	default:
		return nil
	}
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

	commitObj := &models.CommitModel{
		Tree:     extractDataFromCommit(commitData, "tree").(string),
		Parent:   extractDataFromCommit(commitData, "parent").(string),
		Author:   extractDataFromCommit(commitData, "author").(models.Author),
		Commiter: extractDataFromCommit(commitData, "commiter").(models.Commiter),
		Message:  commitData[len(commitData)-2],
		Type:     isTypeMerge(commitData),
	}

	return commitObj, nil
}
