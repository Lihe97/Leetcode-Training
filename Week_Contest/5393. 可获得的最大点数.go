package main

import "fmt"

func maxScore(cardPoints []int, k int) int {

	begin := cardPoints[:k]
	end := cardPoints[len(cardPoints)-k:]
	res := 0
	sum1 := 0
	sum2 := 0

	for _,v := range end{
		sum2 += v
	}
	for _,v := range begin{
		sum1 += v
	}


	for i := 0 ; i < k ; i ++{

		if sum1 >= sum2{
			res += begin[0]
			sum1 -= begin[0]
			sum2 -= end[0]
			begin = begin[1:]
			end = end[1:]

		}else{
			res += end[len(end)-1]
			sum2 -= end[len(end)-1]
			sum1 -= begin[len(begin)-1]
			begin = begin[:len(begin)-1]
			end = end[:len(end)-1]
		}

	}
	return res

}

func main() {
	a := []int{1,2,3,4,5,6,1}
	fmt.Println(maxScore(a,3))
	
}
