package utils

import "strings"

func Truncate(str string, length int) string {
	if len(str) <= length {
		return str
	}
	return strings.Join([]string{str[:length], "…"}, "")
}
