package main

import (
	"fmt"
	"git-stats/cmd"
	"os"
	"path"
)

func main() {
	rootPath, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	gitPath := path.Join(rootPath, ".git")
	cmd.GitStats(gitPath)
}
