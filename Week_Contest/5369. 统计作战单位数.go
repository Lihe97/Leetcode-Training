package main

import "fmt"

func numTeams(rating []int) int {

	res := 0

	for i := 1 ; i < len(rating)-1; i++{
		xb := 0
		xa := 0
		db := 0
		da := 0
		for j := 0 ; j < i ; j ++{
			if rating[j] > rating[i]{
				db ++
			}
			if rating[j] < rating[i]{
				xb ++
			}
		}
		for j := i+1 ; j  < len(rating) ; j++{
			if rating[j] > rating[i]{
				da ++
			}
			if rating[j] < rating[i]{
				xa ++
			}
		}
		res += db * xa + xb * da
	}
	return res

}
func main() {

	a := []int{1,2,3,4}
	fmt.Println(numTeams(a))
	
}
