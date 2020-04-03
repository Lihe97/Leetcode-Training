package main

import "fmt"
//
//func searchInsert(nums []int, target int) int {
//
//	return binSearch(nums,0,len(nums)-1,target)
//}
//func binSearch(nums []int,start,end,target int)int{
//	mid := (start + end)/2
//	if nums[mid] == target{
//		return mid
//	}
//	if end - start == 0 && nums[start] < target{
//		return start + 1
//	}else if end - start == 0 && nums[start] > target{
//		return start
//	}
//	if mid == 0 && nums[mid] > target{
//		return mid
//	}else if nums[mid] > target{
//		return binSearch(nums,0,mid-1,target)
//	}
//	if mid == end && nums[mid] < target {
//		return mid + 1
//	}else if nums[mid] < target{
//		return binSearch(nums,mid+1,end,target)
//	}
//	return 0
//}
func searchInsert(nums []int, target int) int {

	left := 0
	right := len(nums)-1

	for left <= right{
		mid := (left + right)/2
		if nums[mid] == target{
			return mid
		}else if nums[mid] < target{
			left = mid + 1

		}else if nums[mid] > target{
			right = mid - 1
		}
	}
	return left
}
func main() {

	a := []int{1,3,5,7,8}
	fmt.Println(searchInsert(a,6))


	
}
