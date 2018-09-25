package utils

import (
	"fmt"
	"strconv"
)

func getCommitKeys(commits map[int]int) []int {
	var (
		keys []int
	)

	for key := range commits {
		keys = append(keys, key)
	}

	return keys
}

func drawCommitCell(commitValue int, isToday bool) {
	dataToBePrinted := ""

	switch {
	case commitValue > 0 && commitValue < 5:
		dataToBePrinted = Commited0To5Color
	case commitValue >= 5 && commitValue < 10:
		dataToBePrinted = Commited5To10Color
	case commitValue >= 10:
		dataToBePrinted = CommitedMoreThan10Color
	case commitValue >= 100:
		dataToBePrinted = CommitedMoreThan100Color
	}

	if isToday {
		dataToBePrinted = TodaysCellColor
	}

	space := "  %s "

	switch {
	case commitValue >= 10:
		space = " %s "
	case commitValue >= 100:
		space = "%s "
	}

	data := "-"
	if commitValue != 0 {
		data = strconv.Itoa(commitValue)
	}

	dataToBePrinted = dataToBePrinted + space + EndColor
	fmt.Printf(dataToBePrinted, data)
}

func drawTopBottomBoundries() {
	for j := MaxWeeks + 1; j >= 0; j-- {
		fmt.Printf(BoundaryColor, "====")
	}
}

func drawLeftRightBoundries() {
	fmt.Printf(BoundaryColor, " || ")
}

func processCommitCells(graph [7][MaxWeeks]int) {
	for i := 6; i >= 0; i-- {
		if i == 6 {
			drawTopBottomBoundries()
			fmt.Printf("\n")
		}

		drawLeftRightBoundries()

		for j := MaxWeeks - 1; j >= 0; j-- {
			if i == 0 && j == 0 {
				drawCommitCell(graph[i][j], true)
			} else {
				drawCommitCell(graph[i][j], false)
			}
		}

		drawLeftRightBoundries()

		fmt.Printf("\n")

		if i == 0 {
			drawTopBottomBoundries()
			fmt.Printf("\n")
		}
	}
}

// PlotCommits ... Plot the generated commits
func PlotCommits(commits map[int]int) {
	var (
		graph [7][MaxWeeks]int
	)

	keys := getCommitKeys(commits)

	for _, key := range keys {
		week := int(key / 7)
		day := key % 7

		fmt.Println(week, day, commits[key])
		graph[day][week] = commits[key]
	}

	// fmt.Println(graph)
	processCommitCells(graph)
}
