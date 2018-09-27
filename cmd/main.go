package cmd

import (
	"fmt"
	"git-stats/internals/constants"
	"git-stats/internals/git"
	"git-stats/internals/stats"
	"path"
	"strconv"
)

// GitStats ... Run git stats
func GitStats(gitPath string) {
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

	// constants.PrintAvailableColors()

	userEmail, err := git.GetUserEmail(gitPath)
	fmt.Println(constants.YellowText + "Branch: " + currentBranchPath + constants.EndText)
	fmt.Println(constants.DeepBlueText + "User Email: " + userEmail + constants.EndText)

	commitList, commitMap := GenerateCommitList(gitPath, headSHA)
	fmt.Println("\n" + constants.WhiteText + strconv.Itoa(commitList.Len()) + " contributions in " + currentBranchPath + constants.EndText)
	stats.PlotCommits(commitMap)
	stats.ProcessCommits(commitList, userEmail)
}
