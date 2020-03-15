package main

import "fmt"

func missingNumber(nums []int) int {
	l := len(nums)
	left := 0
	right := l-1

	for left <= right{
		mid := (left + right)/2
		if nums[mid] != mid{
			right = mid - 1
		}else{
			left = mid + 1
		}
	}
	return left

}
func main() {

	a := []int{0,1,2,3,5,6,7,8}
	fmt.Println(missingNumber(a))
}
