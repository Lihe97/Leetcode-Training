package main

import (
	"fmt"
	"math"
	"sort"
)

func minDifference(nums []int) int {

	if len(nums) <= 4{
		return 0
	}
	sort.Ints(nums)
	ans := math.MaxInt16

	l,r := 0,len(nums)-1

	ans = min(ans,nums[r]-nums[l+3])
	ans = min(ans,nums[r-1]-nums[l+2])
	ans = min(ans,nums[r-2]-nums[l+1])
	ans = min(ans,nums[r-3]-nums[l])
	return ans
}
func min(a,b int)int{
	if a < b{
		return a
	}else{
		return b
	}
}

func main() {

	fmt.Println(minDifference([]int{1,5,0,10,14}))
	
}
