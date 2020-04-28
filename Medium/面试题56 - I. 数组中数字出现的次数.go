package main

import "fmt"

func singleNumbers(nums []int) []int {

	res := []int{}

	num := 0

	for i := 0 ; i < len(nums) ; i ++{
		num ^= nums[i]
	}
	mask := 1
	for mask & num == 0{
		mask <<= 1
	}
	a , b := 0,0

	for i := 0 ; i < len(nums) ; i++{
		if mask & nums[i] == 0{
			a ^= nums[i]
		}else{
			b ^= nums[i]
		}
	}
	res = append(res, a)
	res = append(res, b)
	return res
}

func main() {

	a := []int{0,0,1,2}
	fmt.Println(singleNumbers(a))
	fmt.Println(1<<1)
}
