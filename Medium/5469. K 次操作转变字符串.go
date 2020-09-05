package main

import (
	"fmt"
)

func canConvertString(s string, t string, k int) bool {

	if len(s) != len(t) {
		return false
	}

	a := make([]int,26)

	for i := 0 ; i < len(s) ; i ++{
		if s[i] != t[i]{
			j := (int(t[i]) - int(s[i]) + 26)%26

			a[j] ++
			if j + (a[j] - 1)*26 > k {
				return false
			}
		}

	}
	return true

}

func main() {

	fmt.Print(canConvertString("bpmaaaljbfdp","djzbvyjrkkbs",115))

}