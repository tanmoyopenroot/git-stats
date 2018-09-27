package stats

import (
	"fmt"
	"git-stats/internals/constants"
	"strconv"
	"time"
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
		dataToBePrinted = constants.Commited0To5Color
	case commitValue >= 5 && commitValue < 10:
		dataToBePrinted = constants.Commited5To10Color
	case commitValue >= 10:
		dataToBePrinted = constants.CommitedMoreThan10Color
	case commitValue >= 100:
		dataToBePrinted = constants.CommitedMoreThan100Color
	}

	if isToday {
		dataToBePrinted = constants.TodaysCellColor
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

	dataToBePrinted = dataToBePrinted + space + constants.EndColor
	fmt.Printf(dataToBePrinted, data)
}

func drawTopBottomBoundries() {
	for j := constants.MaxWeeks + 1; j >= 0; j-- {
		fmt.Printf(constants.BoundaryColor, "----")
	}
}

func drawLeftRightBoundries() {
	fmt.Printf(constants.BoundaryColor, "|| ")
}

func printMonths() {
	var (
		currentTime  time.Time
		pastTime     time.Time
		currentMonth time.Month
		prevMonth    time.Month
	)

	currentTime = time.Now()
	pastTime = currentTime.Add(-constants.MaxWeeks * time.Hour * 24 * 7)

	fmt.Printf(constants.BoundaryColor, "|      ")

	for pastTime.Before(currentTime) {
		currentMonth = pastTime.Month()
		if currentMonth != prevMonth {
			fmt.Printf(constants.MonthColor, currentMonth.String()[:3])
			prevMonth = currentMonth
		} else {
			fmt.Print("    ")
		}
		pastTime = pastTime.Add(time.Hour * 24 * 7)
	}
	fmt.Printf(constants.BoundaryColor, "|")
}

func printWeekdays(day int) {
	switch day {
	case 5:
		fmt.Printf(constants.WeekDayColor, "Mon")
		break
	case 3:
		fmt.Printf(constants.WeekDayColor, "Wed")
		break
	case 1:
		fmt.Printf(constants.WeekDayColor, "Fri")
		break
	default:
		fmt.Printf("   ")
	}
}

func processCommitCells(graph [7][constants.MaxWeeks]int) {
	for i := 6; i >= 0; i-- {
		if i == 6 {
			drawTopBottomBoundries()
			fmt.Printf("\n")
			printMonths()
			fmt.Printf("\n")
		}

		drawLeftRightBoundries()
		printWeekdays(i)

		for j := constants.MaxWeeks - 1; j >= 0; j-- {
			if i == (6-int(time.Now().Weekday())) && j == 0 {
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

func drawIndicators() {
	fmt.Println("   Indicators:")
	fmt.Printf("   Current Day:  "+constants.TodaysCellColor+"  %s "+constants.EndColor, "-")
	fmt.Printf("\n")
	fmt.Printf("   Commits:      "+constants.Commited0To5Color+"  %s "+constants.EndColor, "-")
	fmt.Printf("   "+constants.Commited5To10Color+"  %s "+constants.EndColor, "-")
	fmt.Printf("   "+constants.CommitedMoreThan10Color+"  %s "+constants.EndColor, "-")

	fmt.Printf("\n")
}

// PlotCommits ... Plot the generated commits
func PlotCommits(commits map[int]int) {
	var (
		graph [7][constants.MaxWeeks]int
	)

	keys := getCommitKeys(commits)
	currentTime := time.Now()
	currentDayNumber := int(currentTime.Weekday())
	offset := 6 - currentDayNumber

	for _, key := range keys {
		week := int((key + offset) / 7)
		day := (key + offset) % 7

		graph[day][week] = commits[key]
	}

	// fmt.Println(graph)
	processCommitCells(graph)
	drawIndicators()
}
