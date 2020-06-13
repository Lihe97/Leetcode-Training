package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)

	for i := 0 ; i < len(nums) - 2 ; i++{
		if i > 0 && nums[i] == nums[i-1]{
			continue
		}
		low := i + 1
		high := len(nums) - 1
		for low < high{
			sum := nums[i] + nums[low] + nums[high]
			if sum == 0{

				res = append(res, []int{nums[i], nums[low], nums[high]})
				for low < high && nums[low] == nums[low+1]{
					low ++
				}
				for low < high && nums[high] == nums[high - 1]{
					high --
				}
				low ++
				high --
			}else if sum < 0{
				low ++
			}else if sum > 0{
				high --
			}

		}

	}
	return res

}

func main() {
	a := []int{-1,0,1,2,-1,-4}
	fmt.Println(threeSum(a))
}
