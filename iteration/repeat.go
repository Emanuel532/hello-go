package iteration

import "strings"

func Repeat(str string, repeatCount int) string {
	var repeated strings.Builder
	for range repeatCount {
		repeated.WriteString(str)
	}
	return repeated.String()
}
