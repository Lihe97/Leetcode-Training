package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {

	res := []int{}
	for i := 0 ; i  < len(nums) ; i ++{

		for nums[nums[i] - 1] != nums[i]{
			swap(&nums[i],&nums[nums[i]-1])
		}
	}
	for i := 0 ; i < len(nums) ; i ++{
		if nums[i] != i+1{
			res = append(res, i+1)
		}
	}
	return res

}
func swap(a,b *int){
	*a = *a ^ *b
	*b = *a ^ *b
	*a = *a ^ *b
}


func main() {

	a := []int{4,3,2,7,8,2,3,1}
	fmt.Println(findDisappearedNumbers(a))
	
}
