package main

import "fmt"

func t2(opr [][3]int,nums []int,m int)  {

	dp := make([]int,600000)
	dp[0] = 1
	dp[1] = 1
	for i := 2 ; i < len(dp) ; i ++{
		dp[i] = dp[i-1] * i
	}

	for i := 0 ; i < m ; i++{
		for j := opr[i][0]-1 ; j <= opr[i][1] - 1 ; j ++{

			temp := j + 1 - opr[i][0] + opr[i][2]

			nums[j] += dp[temp]/(dp[temp-opr[i][2]]*dp[opr[i][2]])
			//fmt.Println(nums,i,j,opr[i],temp)
		}
	}
	for i := 0 ; i < len(nums) ; i ++{
		fmt.Print(nums[i]%998244353)
		fmt.Print(" ")
	}

}

func main() {


	var n int
	var m int
	fmt.Scan(&n)
	fmt.Scan(&m)
	nums := make([]int,n)
	opr := make([][3]int,m)

	for i := 0 ; i < m ; i ++{
		for j := 0 ; j < 3; j ++{
			fmt.Scan(&opr[i][j])
		}
	}
	t2(opr,nums,m)

}