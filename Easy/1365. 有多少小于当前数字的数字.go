package main

import (
	"fmt"
	"sort"
)

func smallerNumbersThanCurrent(nums []int) []int {

	var temp = make([]int, len(nums))
	mp := map[int]int{}
	copy(temp,nums)
	sort.Ints(nums)
	mp[nums[0]] = 0
	for i := 1 ; i < len(nums) ; i++{
		for j := i -1 ; j >= 0 ; j --{
			if nums[i] > nums[j]{
				mp[nums[i]] = j + 1
				break
			}
		}
	}
	res := []int{}
	for i := 0 ; i < len(temp) ; i++{
		res = append(res, mp[temp[i]])
	}
	return res
}
func main() {

	a := []int{6,5,4,8}
	fmt.Println("??")
	fmt.Println(smallerNumbersThanCurrent(a))
	
}
