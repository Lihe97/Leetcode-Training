package main

import "fmt"

func majorityElement(nums []int) int {

	//num := 0
	a := map[int]int{}

	for i := 0 ;i < len(nums);i ++{
		a[nums[i]] ++
		if a[nums[i]] > len(nums)/2{
			return nums[i]
		}
	}
	//
	//fmt.Println(a)
	return 0


}
func main() {
	b := []int{5,5,6,6,6}
	fmt.Println(majorityElement(b))
	
}
