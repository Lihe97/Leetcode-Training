package main

import (
	"fmt"
	"sort"
)

func dominantIndex(nums []int) int {
	if len(nums) ==1 {
		return -1
	}
	a := map[int]int{}
	for x,y := range  nums{
		a[y] =x
	}
	sort.Ints(nums)
	if nums[len(nums)-1] >= 2 * nums[len(nums)-2]{
		return a[nums[len(nums)-1]]
	}
	return -1
}
func main() {

	b := []int{1,2,3,4}
	fmt.Println(dominantIndex(b))
	
}
