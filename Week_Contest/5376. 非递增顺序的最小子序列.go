package main

import (
	"fmt"
	"sort"
)

func minSubsequence(nums []int) []int {

	sum := 0
	for i := 0 ; i < len(nums)  ; i ++{
		sum += nums[i]
	}
	sort.Ints(nums)
	temp := 0
	res := []int{}
	for j := len(nums)-1 ;j >= 0 ; j--{
		temp += nums[j]
		res = append(res, nums[j])
		if temp > sum/2{
			break
		}
	}
	return res

}

func main() {

	a := []int{4,4,6,7,7}
	fmt.Println(minSubsequence(a))
}
