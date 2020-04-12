package main

import (
	"fmt"
	"strconv"
)

func calPoints(ops []string) int {

	stack := []int{}
	for i := 0 ; i < len(ops) ; i ++{

		if ops[i] == "+"{
			stack = append(stack,stack[len(stack)-1]+stack[len(stack)-2])
		}else if ops[i] == "D"{
			stack = append(stack, 2*stack[len(stack)-1])
		}else if ops[i] == "C"{
			stack = stack[:len(stack)-1]
		}else{
			temp,_ := strconv.Atoi(ops[i])
			stack = append(stack, temp)

		}
	}


	res := 0
	for _,a := range stack{
		res += a
	}
	return res
}

func main() {

	a := []string{"5","-2","4","C","D","9","+","+"}
	fmt.Println(calPoints(a))
	
}
