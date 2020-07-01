package main

import (
	"fmt"
)

func canArrange(arr []int, k int) bool {

	temp := []int{}

	for i := 0 ; i < len(arr); i ++{
		temp = append(temp,arr[i]%k)
	}

}

func main() {
	fmt.Println(canArrange([]int{-4,-7,5,2,9,1,10,4,-8,-3},3))

	
}
