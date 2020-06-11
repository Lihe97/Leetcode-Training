package main

import "fmt"

func dailyTemperatures(T []int) []int {

	res := make([]int,len(T))

	if len(T) == 0{
		return res
	}

	stack := []int{}
	stack = append(stack,0)

	for i := 1 ; i < len(T) ; i ++{

		for len(stack) != 0 && T[i] > T[stack[len(stack) - 1]]{

			res[stack[len(stack) - 1]] = i - stack[len(stack) - 1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack,i)

	}
	return res

}

func main() {

	a := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(a))
	
}
