package main

import "fmt"

func singleNumber(nums []int) []int {

	res := []int{}
	temp := 0


	for i := 0 ; i < len(nums); i++{
		temp ^= nums[i]
	}

	mask := 1
	for temp & mask == 0{
		mask <<= 1
	}


	a,b := 0,0

	for i := 0 ; i < len(nums) ; i++{
		if nums[i] & mask == 0{
			a ^= nums[i]
		}else{
			b ^= nums[i]
		}
	}

	res = append(res,a)
	res = append(res,b)
	return res

}

func main() {

	a := []int{1,2,1,3,2,5}
	fmt.Println(singleNumber(a))
	
}
