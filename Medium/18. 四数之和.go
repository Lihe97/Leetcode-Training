package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	if len(nums) < 4 {
		return nil
	}

	for i := 0 ; i < len(nums) - 2 ; i++{
		if i > 0 && nums[i]==nums[i-1]{
			continue
		}
		if target > 0 && nums[i] > target{
			break
		}
		for j := i + 1 ; j < len(nums) - 1; j++{
			low := j + 1
			high := len(nums) - 1
			if nums[i]+nums[j]+nums[low] > target && nums[low] > 0 { // 排除不可能的情况
				continue
			}
			if j > i+1 && nums[j] == nums[j-1] { // 去重
				continue
			}
			for low < high{
				if low > j+1 && nums[low] == nums[low-1] { // 去重
					low++
					continue
				}

				if high < len(nums) - 2 && nums[high] == nums[high+1] { // 去重
					high--
					continue
				}
				sum := nums[i] + nums[low] + nums[high] + nums[j]
				if sum == target{
					res = append(res, []int{nums[i], nums[low], nums[high],nums[j]})
					for low < high && nums[low] == nums[low+1]{
						low ++
				}
					for low < high && nums[high] == nums[high - 1]{
						high --
				}
					low ++
					high --
				}else if sum < target{
					low ++
					continue
				}else if sum > target{
					high --
					continue
				}
			}
		}
	}
	return res

}

func main() {
	a := []int{0,0,0,0}

	fmt.Println(fourSum(a,1))
	
}
