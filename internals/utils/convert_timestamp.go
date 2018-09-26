package utils

import (
	"fmt"
	"strconv"
	"time"
)

// ConvertTimeStamp ... Convert timestamp to type time.Time
func ConvertTimeStamp(timeStamp string) time.Time {
	i, err := strconv.ParseInt(timeStamp, 10, 64)
	if err != nil {
		fmt.Println(err)
		panic("Cannot parse timestamp")
	}

	return time.Unix(i, 0)
}
