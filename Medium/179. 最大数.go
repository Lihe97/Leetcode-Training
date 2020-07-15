package main

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string{

	res := ""
	sort.Slice(nums, func(i, j int) bool {
		a := strconv.Itoa(nums[i])
		b := strconv.Itoa(nums[j])
		return a + b > b + a
	})

	for i := 0 ; i < len(nums) ; i ++{
		res += strconv.Itoa(nums[i])
	}
	if nums[0] == 0{
		return "0"
	}
	return res
}

func main() {
	
}
