package utils

import (
	"time"
)

func getBeginningOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	startOfDay := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	return startOfDay
}

// CountDaysFromNow ... Count number of days from now
func CountDaysFromNow(commitDate time.Time) int {
	days := 0
	now := getBeginningOfDay(time.Now())
	for commitDate.Before(now) {
		commitDate = commitDate.Add(time.Hour * 24)
		days++
	}
	return days
}
