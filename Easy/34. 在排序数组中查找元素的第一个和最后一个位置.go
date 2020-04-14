package main

import "fmt"

//func searchRange(nums []int, target int) []int {
//
//	i := 0
//	j := len(nums)-1
//	res := []int{}
//	if len(nums) == 1{
//		if nums[0] == target{
//			  res = []int{0,0}
//			return res
//		}else{
//			res = []int{-1,-1}
//			return  res
//		}
//	}
//
//
//	for i <= j && (nums[i] != target || nums[j] != target){
//		if nums[i] != target{
//			i ++
//		}
//		if nums[j] != target{
//			j --
//		}
//	}
//
//	if i > j {
//		res = []int{-1,-1}
//	}else{
//		res = []int{i,j}
//	}
//
//	return res
//
//}

func searchRange(nums []int, target int) []int {

	res := []int{-1,-1}
	if len(nums) == 0{
		return res
	}
	left := 0
	right := len(nums)-1

	for left < right{

		mid := (left + right) / 2
		if nums[mid] >= target{
			right = mid
		}else{
			left = mid + 1
		}
	}
	if nums[left] != target{
		return res
	}
	res[0] = left

	left = 0
	right = len(nums)

	for left < right{
		mid := (left + right)/2

		if nums[mid] <= target{
			left = mid + 1
		}else{
			right = mid
		}

	}
	res[1] = left-1


	return res

}

func main() {

	a := []int{1,2,2,3}
	fmt.Println(searchRange(a,4))
	
}
