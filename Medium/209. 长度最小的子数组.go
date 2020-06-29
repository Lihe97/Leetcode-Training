package main

import "math"

func minSubArrayLen(s int, nums []int) int {

	left,right := 0,0
	sum := 0
	res := math.MaxInt16

	for right < len(nums){
		sum += nums[right]
		for sum >= s{
			if right - left + 1 < res{
				res= right - left + 1
			}
			sum -= nums[left]
			left ++
		}
		right ++
	}
	if res == math.MaxInt16{
		return 0
	}

	return res
}

func main() {
	
}
