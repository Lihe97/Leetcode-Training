package main

import "fmt"

func findKthLargest(nums []int, k int) int {

	temp := quickSort(nums,0,len(nums)-1,k)
	return temp[len(nums) - k ]
}
func quickSort(nums []int,left,right,k int) []int{

	if left < right{
		pivot := partition(nums,left,right)
		if pivot == len(nums) - k{
			return nums
		}
		quickSort(nums,left,pivot-1,k)
		quickSort(nums,pivot+1,right,k)
	}
	return nums
}
func partition(nums []int,left ,right int) int{
	pivot := nums[left]
	for left < right{

		for left < right && nums[right] >= pivot{
			right --
		}
		nums[left],nums[right] = nums[right],nums[left]
		for left < right && nums[left] <= pivot{
			left ++
		}
		nums[left],nums[right] = nums[right],nums[left]

	}
	return left
}

func main() {

	a := []int{1,3,2,6,4,7}
	fmt.Println(findKthLargest(a,2))
	
}
