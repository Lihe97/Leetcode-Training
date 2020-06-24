package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {

	max := math.MaxInt8
	res := 0
	sort.Ints(nums)

	for i := 0 ; i < len(nums) - 2 ; i ++{

		left := i + 1
		right := len(nums) - 1

		for left < right{
			sum := nums[i] + nums[left] + nums[right]
			td := sum - target
			if td < 0 {
				td = - td
			}
			if td == 0{
				return sum
			}
			if td < max{
				res = sum
				max = td
			}
			if sum  > target{
				right --
			}else{
				left ++
			}
		}
	}
	return res

}

func main() {
	
}
