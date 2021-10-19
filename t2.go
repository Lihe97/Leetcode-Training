package main

import (
	"strings"
)

func removeOccurrences(s string, part string) string {

	find := strings.Contains(s, part)
	for find {
		s = strings.Replace(s, part, "", 1)
		find = strings.Contains(s, part)
	}
	return s
}
func main() {




}