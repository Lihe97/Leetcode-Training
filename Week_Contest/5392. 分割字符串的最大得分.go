package main

import "fmt"

func maxScore(s string) int {


	max := 0

	for i := 0 ; i < len(s)-1 ; i++{
		count0 := 0
		count1 := 0
		for j := 0 ; j <= i ; j ++{
			if s[j] == '0'{
				count0++
			}
		}
		for j := i + 1; j < len(s) ; j ++{
			if s[j] == '1'{
				count1 ++
			}
		}
		if count0 + count1 > max{
			max = count0 + count1
		}

	}
	return max
}

func main() {

	fmt.Println(maxScore("00111"))
}
