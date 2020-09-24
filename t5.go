package main

import "fmt"

func t1(nums [][2]int)  {

	m := -1

	for i := 0 ; i < len(nums) ; i ++{
		if nums[i][1] > m{
			m = nums[i][1]
		}
	}
	flag := make([]bool,m+1)
	for i := 0 ; i < len(nums) ; i ++{
		for j := nums[i][0] ; j <= nums[i][1] - 1; j ++{
			flag[j] = true
		}
	}
	sum := 0
	for i := 0 ; i < len(flag) ; i ++{
		if flag[i]{
			sum ++
		}
	}
	fmt.Print(sum)

}

func main() {

	var n int

	fmt.Scan(&n)
	nums := make([][2]int,n)

	for i := 0 ; i < n ; i ++{
		fmt.Scan(&nums[i][0])
		fmt.Scan(&nums[i][1])
	}
	t1(nums)

}
