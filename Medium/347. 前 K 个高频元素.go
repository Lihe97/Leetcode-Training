package main

import "sort"

func topKFrequent(nums []int, k int) []int {

	mp := map[int]int{}

	temp := []int{}

	for i := 0 ; i  < len(nums) ; i ++{
		mp[nums[i]] ++
	}
	for a,_ := range mp{
		temp = append(temp,a)
	}
	sort.Slice(temp, func(i, j int) bool {
		return mp[temp[i]] < mp[temp[j]]
	})
	res := []int{}
	for j := len(temp) -1 ; j >= len(temp) - k ; j --{
		res = append(res,temp[j])
	}
	return res
}
func main() {
	
}
