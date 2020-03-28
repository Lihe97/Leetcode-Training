package main

import (
	"fmt"
	"sort"
)

func singleNumber(nums []int) int {

	sort.Ints(nums)
	if nums[0] != nums[1] {
		return nums[0]
	}
	if nums[len(nums)-1] != nums[len(nums)-2]{
		return nums[len(nums)-1]
	}
	for i := 2 ; i < len(nums)-2 ;i++{
		if nums[i] != nums[i-1] && nums[i] != nums[i+1]{
			return nums[i]
		}
	}
	return 0

}

func main() {
	a := []int{3,3,1,4,4,5,5,5}
	fmt.Println(singleNumber(a))
	
}
