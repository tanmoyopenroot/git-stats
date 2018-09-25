package git

import (
	"bytes"
	"compress/zlib"
	"git-stats/internals/models"
	"git-stats/internals/utils"
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

func extractTimeStamp(data string) string {
	splitData := strings.Split(data, " ")
	return splitData[len(splitData)-2]
}

func extractCommiter(data string) models.Commiter {
	return models.Commiter{
		Name:      utils.ExtractString(data, "committer", "<"),
		Email:     utils.ExtractString(data, "<", ">"),
		TimeStamp: extractTimeStamp(data),
	}
}

func extractAuthor(data string) models.Author {
	return models.Author{
		Name:      utils.ExtractString(data, "author", "<"),
		Email:     utils.ExtractString(data, "<", ">"),
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
				return utils.ExtractString(data[i], "tree", "")
			case "parent":
				return utils.ExtractString(data[i], "parent", "")
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
		Tree:      extractDataFromCommit(commitData, "tree").(string),
		Parent:    extractDataFromCommit(commitData, "parent").(string),
		Author:    extractDataFromCommit(commitData, "author").(models.Author),
		Commiter:  extractDataFromCommit(commitData, "commiter").(models.Commiter),
		Message:   commitData[len(commitData)-2],
		TypeMerge: isTypeMerge(commitData),
	}

	return commitObj, nil
}
