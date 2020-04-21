package main

import (
	"fmt"
)

func numberOfSubarrays(nums []int, k int) int {

	temp := []int{}
	res := 0
	temp = append(temp, -1)
	for i := 0 ; i < len(nums);  i ++{
		if nums[i] % 2 == 1{
			temp = append(temp, i)
		}
	}
	temp = append(temp, len(nums))

	for i := 1; i + k < len(temp) ;i++ {
		res += (temp[i] - temp[i - 1]) * (temp[i + k] - temp[i + k - 1]);
	}

	return res

}

func main() {

	a := []int{1,1,2,1,1}
	fmt.Println(numberOfSubarrays(a,3))
}
