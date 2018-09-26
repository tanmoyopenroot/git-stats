package stats

import (
	"container/list"
	"fmt"
	"git-stats/internals/constants"
	"git-stats/internals/models"
	"git-stats/internals/utils"
	"strconv"
	"strings"
)

func displayCommits(commitList *list.List) {
	if commitList.Len() > 0 {
		fmt.Println("\n" + constants.WhiteText + "COMMITS(Latest " + strconv.Itoa(constants.ShowMaxCommits) + ")" + constants.EndText)
		fmt.Print("\n")
	}

	for commit := commitList.Front(); commit != nil; commit = commit.Next() {
		data := commit.Value.(*models.CommitModel)
		authorCommitTime := strings.Split(utils.ConvertTimeStamp(data.Author.TimeStamp).String(), " ")[0]
		fmt.Println("Message: " + constants.YellowText + data.Message + constants.EndText)
		fmt.Println(constants.GreyText + data.Author.Name + " authored on " + authorCommitTime + constants.EndText)
		fmt.Print("\n")
	}
}

func displayMerges(mergeList *list.List) {
	if mergeList.Len() > 0 {
		fmt.Println("\n" + constants.WhiteText + "MERGES(Latest " + strconv.Itoa(constants.ShowMaxMerges) + ")" + constants.EndText)
		fmt.Print("\n")
	}

	for commit := mergeList.Front(); commit != nil; commit = commit.Next() {
		data := commit.Value.(*models.CommitModel)
		authorCommitTime := strings.Split(utils.ConvertTimeStamp(data.Author.TimeStamp).String(), " ")[0]
		fmt.Println("Message: " + constants.YellowText + data.Message + constants.EndText)
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

	if commitList.Len() > 0 {
		fmt.Println("\n" + constants.GreenText + "CONTRIBUTION ACTIVITY" + constants.EndText)
	}

	displayCommits(lastestCommitList)
	displayMerges(lastestMergeList)
}
