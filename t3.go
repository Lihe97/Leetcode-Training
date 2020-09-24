package main

import (
	"fmt"
	"sort"
)

func t3(nums []int)  {

	sort.Ints(nums)
	fmt.Println(nums[len(nums)/2])
}

func main() {

	var n int
	fmt.Scan(&n)
	nums := make([]int,n)
	for i := 0 ; i < n ; i ++{
		fmt.Scan(&nums[i])
	}

	t3(nums)



}