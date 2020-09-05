package main

import "fmt"

func stoneGameV(stoneValue []int) int {

	if len(stoneValue) == 1{
		return 0
	}
	sum := 0

	res := 0


	for len(stoneValue) != 1{
		sum = 0
		for i := 0 ; i < len(stoneValue) ; i ++{
			sum += stoneValue[i]
		}
		temp := 0
		for i := 0 ; i < len(stoneValue) - 1 ; i ++{
			temp += stoneValue[i]
			if temp + stoneValue[i+1] > sum/2{
				t := sum - temp
				if temp > t{
					res += t
					stoneValue = stoneValue[i+1:]
				}else{
					res += temp
					stoneValue = stoneValue[:i+1]
				}
			}
		}
		fmt.Println(stoneValue,sum,res)
	}

	return res

}
func main() {

	fmt.Println(stoneGameV([]int{9,8,2,4,6,3,5,1,7}))



}