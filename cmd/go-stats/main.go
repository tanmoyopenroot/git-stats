package main

import (
	"fmt"
	"git-stats/internals/git"
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
	currentBranchPath, err := git.CurrentBranchPath(gitPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	headSHA, err := git.GetSHA(path.Join(gitPath, currentBranchPath))
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = GenerateCommitList(gitPath, headSHA)
	// for commit := commitList.Front(); commit != nil; commit = commit.Next() {
	// 	fmt.Println(commit.Value)
	// }
}
