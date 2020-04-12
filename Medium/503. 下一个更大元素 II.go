package main

import (
	"fmt"
)

func nextGreaterElements(nums []int) []int {
	temp := []int{}
	res := []int{}
	mp := map[int]int{}

	for i := 0 ; i < len(nums) ; i ++{

		for len(temp) != 0 && nums[i] > nums[temp[len(temp)-1]]{
			mp[temp[len(temp)-1]] = i
			temp = temp[0:len(temp)-1]
		}
		temp = append(temp, i)
	}
	for i := 0 ; i < len(nums) ; i ++{

		for len(temp) != 0 && nums[i] > nums[temp[len(temp)-1]]{
			mp[temp[len(temp)-1]] = i
			temp = temp[0:len(temp)-1]
		}
	}
	for i := 0 ; i < len(temp); i ++{
		mp[temp[i]] = -1
	}

	for i := 0 ; i < len(nums) ; i ++{
		if mp[i] != -1{
			res = append(res, nums[mp[i]])
		}else{
			res = append(res, mp[i])
		}

	}



	return res
}
func main() {
	a := []int{1,2,3,2,1,3}
	fmt.Println(nextGreaterElements(a))


	
}
