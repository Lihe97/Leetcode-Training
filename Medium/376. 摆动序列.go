package main

import "fmt"

func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	begin, up, down := 0, 1, 2
	state := begin
	maxLen := 1
	for i := 1; i <len(nums); i++{
		switch state {
		case begin:
			if nums[i]>nums[i-1]{
				state = up
				maxLen++
			}else if nums[i] < nums[i-1]{
				state = down
				maxLen++
			}
		case up:
			if nums[i] < nums[i-1] {
				state = down
				maxLen++
			}
		case down :
			if nums[i] > nums[i-1] {
				state = up
				maxLen++
			}
		}
	}
	return  maxLen
}
func main() {
	a:= []int{1,3,2,5,1,9}
	fmt.Println(wiggleMaxLength(a))
}
