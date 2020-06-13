package main

import (
	"fmt"
	"sort"
)

func minSumOfLengths(arr []int, target int) int {

	res := []int{}


	left,right := 0,0
	cnt := 0
	cur := 0
	temp := -1

	for right < len(arr){
		if arr[right] + cur < target{
			cur += arr[right]
			right ++
			cnt ++
		}else if arr[right] + cur == target{
			cnt ++
			//fmt.Println("前",left,right,temp,cnt)

			if left < temp && cnt < res[len(res) - 1]{
				res[len(res) - 1] = cnt
				fmt.Println("更新",cnt)
			}else if left > temp + 1{
				temp = right
				res = append(res,cnt)
				fmt.Println("插入",cnt)
			}
			//fmt.Println("后",left,right,temp,cnt)

			cur -= arr[left]
			cnt -= 2
			left = left + 1
		}else if arr[right] + cur > target{
			cur -= arr[left]
			left ++
			cnt --
		}
		fmt.Println(left,right)

	}

	fmt.Println(res)
	sort.Ints(res)
	if len(res) >= 2{
		return res[0] + res[1]
	}else{
		return -1

	}

}

func main() {

	a := []int{34,6,8,1,1,2,45,2,2,1,1,1,50,1,1}

	fmt.Println(minSumOfLengths(a,52))
	
}
