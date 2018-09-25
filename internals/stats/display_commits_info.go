package stats

import (
	"container/list"
	"fmt"
	"git-stats/internals/constants"
	"git-stats/internals/models"
	"strconv"
	"strings"
	"time"
)

func displayCommits(commitList *list.List) {
	if commitList.Len() > 0 {
		fmt.Println("\n" + constants.GreenText + "Last " + strconv.Itoa(constants.ShowMaxCommits) + " commits from the current branch:" + constants.EndText)
		fmt.Print("\n")
	}

	for commit := commitList.Front(); commit != nil; commit = commit.Next() {
		data := commit.Value.(*models.CommitModel)
		i, err := strconv.ParseInt(data.Author.TimeStamp, 10, 64)
		if err != nil {
			fmt.Println(err)
			break
		}
		authorCommitTime := strings.Split(time.Unix(i, 0).String(), " ")[0]
		fmt.Println(constants.YellowText + data.Message + constants.EndText)
		fmt.Println(constants.GreyText + data.Author.Name + " authored on " + authorCommitTime + constants.EndText)
		fmt.Print("\n")
	}
}

func displayMerges(mergeList *list.List) {
	if mergeList.Len() > 0 {
		fmt.Println("\n" + constants.GreenText + "Last " + strconv.Itoa(constants.ShowMaxMerges) + " merges from the current branch:" + constants.EndText)
		fmt.Print("\n")
	}

	for commit := mergeList.Front(); commit != nil; commit = commit.Next() {
		data := commit.Value.(*models.CommitModel)
		i, err := strconv.ParseInt(data.Author.TimeStamp, 10, 64)
		if err != nil {
			fmt.Println(err)
			break
		}
		authorCommitTime := strings.Split(time.Unix(i, 0).String(), " ")[0]
		fmt.Println(constants.YellowText + data.Message + constants.EndText)
		fmt.Println(constants.GreyText + data.Author.Name + " authored on " + authorCommitTime + constants.EndText)
		fmt.Print("\n")
	}
}

// ProcessCommits ... Show latest n commits and merges
func ProcessCommits(commitList *list.List, userEmail string) {
	numCommits := 0
	numMerges := 0

	lastestCommitList := list.New()
	lastestMergeList := list.New()

	for commit := commitList.Front(); commit != nil; commit = commit.Next() {
		data := commit.Value.(*models.CommitModel)
		if data.Author.Email == userEmail {
			if data.TypeMerge && numMerges < constants.ShowMaxMerges {
				lastestMergeList.PushBack(data)
				numMerges++
			} else if !data.TypeMerge && numCommits < constants.ShowMaxCommits {
				lastestCommitList.PushBack(data)
				numCommits++
			} else {
				break
			}
		}
	}

	displayCommits(lastestCommitList)
	displayMerges(lastestMergeList)
}
