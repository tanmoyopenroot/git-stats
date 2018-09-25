package main

import (
	"fmt"
	"git-stats/internals/git"
	"git-stats/internals/stats"
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

	userEmail, err := git.GetUserEmail(gitPath)
	fmt.Println(userEmail)
	commitList, commitMap := GenerateCommitList(gitPath, headSHA)
	// constants.PrintAvailableColors()
	stats.PlotCommits(commitMap)
	stats.ProcessCommits(commitList, userEmail)
}
