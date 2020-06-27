package main

import "fmt"

func longestSubarray(nums []int) int {

	res := [][]int{}

	sum1 := 0
	sum0 := 0

	max := 0
	mm := 0
	f := false

	for i := 0 ; i < len(nums) ; i ++{

		if nums[i] == 1{
			if sum0 == 0 {
				sum1 ++
			}else{
				res = append(res,[]int{0,sum0})
				f = true
				sum0 = 0
				sum1 ++
			}
		}else if nums[i] == 0{
			if sum1 == 0{
				sum0 ++
			}else{
				if sum1 > mm{
					mm = sum1
				}
				res = append(res,[]int{1,sum1})
				sum1 = 0
				sum0 ++
			}
		}
	}
	if sum0 != 0{
		res = append(res,[]int{0,sum0})
		f = true
	}else{
		if sum1 > mm{
			mm = sum1
		}
		res = append(res,[]int{1,sum1})
	}
	flag := false
	for i := 0 ; i < len(res) ; i ++{
		temp := 0
		if res[i][0] == 0 && res[i][1] == 1{
			flag = true
			if i < len(res)-1{
				temp += res[i+1][1]
			}
			if i > 0{
				temp += res[i-1][1]
			}
		}
		if res[i][0] == 1{
			temp = res[i][1]
		}

		if temp > max{
			max = temp
		}
	}
	//fmt.Println(flag)
	if flag == false && mm != 0 && f {
		return mm
	}
	if flag == false && mm != 0 && !f {
		return mm-1
	}
	return max

}

func main() {
	fmt.Println(longestSubarray([]int{1,0,0,1,1,1,1,1,0,0,0,0,0,0,1,1,1,0,1,1,1,1,1,0,1,1,1,1,1,1,1,1,0,1,1,0,0,0,1,1,0,0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,0,0,0,0,1,1,1,1,0,1,1}))

	
}
