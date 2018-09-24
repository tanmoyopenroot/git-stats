package utils

import (
	"fmt"
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

func printCell(val int, today bool) {
	escape := "\033[0;37;30m"
	switch {
	case val > 0 && val < 5:
		escape = "\033[1;30;47m"
	case val >= 5 && val < 10:
		escape = "\033[1;30;43m"
	case val >= 10:
		escape = "\033[1;30;42m"
	}

	if today {
		escape = "\033[1;37;45m"
	}

	if val == 0 {
		fmt.Printf(escape + "  - " + "\033[0m")
		return
	}

	str := "  %d "
	switch {
	case val >= 10:
		str = " %d "
	case val >= 100:
		str = "%d "
	}

	fmt.Printf(escape+str+"\033[0m", val)
}

func printCells(graph [7][maxWeeks]int) {
	for i := 6; i >= 0; i-- {
		for j := maxWeeks - 1; j >= 0; j-- {
			if i == 0 && j == 0 {
				printCell(graph[i][j], true)
			} else {
				printCell(graph[i][j], false)
			}
		}
		fmt.Printf("\n")
	}
}

// PlotCommits ... Plot the generated commits
func PlotCommits(commits map[int]int) {
	var (
		graph [7][maxWeeks]int
	)

	keys := getCommitKeys(commits)

	for _, key := range keys {
		week := int(key / 7)
		day := key % 7

		fmt.Println(week, day, commits[key])
		graph[day][week] = commits[key]
	}

	// fmt.Println(graph)
	printCells(graph)
}
