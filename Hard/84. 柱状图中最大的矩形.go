package main

import (
	"fmt"
)

func largestRectangleArea(heights []int) int {


	stack := []int{}
	stack = append(stack, -1)
	max := 0
	for i := 0 ; i < len(heights) ; i ++{

		for stack[len(stack)-1] != -1 && heights[stack[len(stack)-1]] >= heights[i]{
			temp := heights[stack[len(stack)-1]]  * ( i - stack[len(stack)-2] - 1)
			stack = stack[0:len(stack)-1]
			if temp > max {
				max = temp
			}

		}
		stack = append(stack, i)

	}
	for stack[len(stack) - 1] != -1{
		temp := heights[stack[len(stack)-1]]  * ( len(heights) - stack[len(stack)-2] - 1)
		stack = stack[0:len(stack)-1]

		if temp > max {
			max = temp
		}

	}

	return max
}

func main() {
	a := []int{2,1,5,6,2,3}
	fmt.Println(largestRectangleArea(a))
}
