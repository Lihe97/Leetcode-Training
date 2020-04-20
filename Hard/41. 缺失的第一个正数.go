package main

import "fmt"

func firstMissingPositive(nums []int) int {

	for i := 0 ; i < len(nums) ; i ++{

		for nums[i]>=1 && nums[i] <= len(nums) && nums[nums[i] - 1] != nums[i]{
			swap(&nums[i],&nums[nums[i]-1])
		}
	}
	for i := 0 ; i < len(nums) ; i ++{
		if nums[i] != i + 1{
			return i + 1
		}
	}

	return len(nums) + 1

}
func swap(a,b *int){
	*a = *a ^ *b
	*b = *a ^ *b
	*a = *a ^ *b
}

func main() {
	a := []int{1,1}
	fmt.Println(firstMissingPositive(a))
	
}
