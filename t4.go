package main

import (
	"fmt"
	"math"
)

func t4(n int)  {
	nums := make([]bool,n)
	for i := 4 ; i < n ; i ++{

		for j := 2 ; j <= int(math.Sqrt(float64(i))) ; j++{
			if i % j == 0{
				nums[i] = true
				break
			}
		}
	}
	temp := []int{}
	for i := 2 ; i < len(nums) ; i ++{
		if !nums[i] {
			temp = append(temp,i)
		}
	}
	left := 0
	right := len(temp)-1
	for temp[left] + temp[right] != n{
		sum := temp[left] + temp[right]
		if sum > n{
			right--
		}else{
			left ++
		}
	}
	fmt.Print(temp[left],temp[right])



}

func main()  {

	var n int
	fmt.Scan(&n)
	t4(n)

}