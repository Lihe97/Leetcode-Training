package main

import (
	"fmt"
	"math"
)

func sumFourDivisors(nums []int) int {
	res := 0
	temp := 0
	for i := 0 ; i < len(nums) ; i++{
		count := 0
		//fmt.Println(nums[i],int(math.Floor(math.Sqrt(float64(nums[i])))))
		for j := 2 ; j <= int(math.Floor(math.Sqrt(float64(nums[i])))); j++{
			if nums[i] % j == 0{

				if j*j == nums[i]{
					break
				}
				count ++
				if count > 1{
					break
				}
				temp = j
			}
		}
		if count == 1{
			res += 1
			res += nums[i]
			res += temp
			res += nums[i]/temp

		}
		//fmt.Println("res",res)
	}
	return res

}
func main() {
	a := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14}
	fmt.Println(sumFourDivisors(a))
}
