package main

import (
	"container/list"
	"fmt"
	"git-stats/internals/git"
	"git-stats/internals/models"
	"git-stats/internals/utils"
	"path"
)

// GenerateCommitList ... Generate list of commits of the current branch
func GenerateCommitList(gitPath string, SHA string) (*list.List, map[int]int) {
	var (
		dataPath string
		data     *models.CommitModel
		err      error
	)

	commitList := list.New()
	commitMap := make(map[int]int)

	for {
		dataPath = path.Join(gitPath, "objects", SHA[:2], SHA[2:])
		data, err = git.ReadObject(dataPath)
		if err != nil {
			fmt.Println(err)
			break
		}

		commitList.PushBack(data)
		daysAgo := utils.CountDaysFromNow(utils.ConvertTimeStamp(data.Author.TimeStamp))
		commitMap[daysAgo]++
		if len(data.Parent) == 0 {
			break
		}
		SHA = data.Parent
	}

	// fmt.Println(commitMap)
	return commitList, commitMap
}
