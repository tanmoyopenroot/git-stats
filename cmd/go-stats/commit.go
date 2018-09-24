package main

import (
	"container/list"
	"fmt"
	"git-stats/internals/git"
	"git-stats/internals/models"
	"git-stats/internals/utils"
	"path"
	"strconv"
	"time"
)

// GenerateCommitList ... Generate list of commits of the current branch
func GenerateCommitList(gitPath string, SHA string) *list.List {
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

		i, err := strconv.ParseInt(data.Author.TimeStamp, 10, 64)
		if err != nil {
			fmt.Println(err)
			break
		}

		daysAgo := utils.CountDaysFromNow(time.Unix(i, 0))
		commitMap[daysAgo]++
		fmt.Println(daysAgo)
		SHA = data.Parent
	}

	fmt.Println(commitMap)

	return commitList
}
