package utils

import "strings"

// ExtractString ... Extract a part of the given data
func ExtractString(data string, start string, end string) string {
	s := 0
	e := len(data)

	if start != "" {
		s = strings.Index(data, start)
		s += len(start)
	}

	if end != "" {
		e = strings.Index(data, end)
	}

	return strings.TrimSpace(data[s:e])
}
