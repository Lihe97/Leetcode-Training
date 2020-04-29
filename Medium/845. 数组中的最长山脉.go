package main

import "fmt"

func longestMountain(A []int) int {

	if len(A) < 3{
		return 0
	}
	dp := make([]int,len(A))
	m := make([]int,len(A))

	dp[0],m[0] = 0,0
	max := 0

	for i := 1; i < len(A) ; i ++{

		if A[i] > A[i-1]{
			m[i] = m[i-1] + 1
		}else{
			m[i] = 0
		}

		if A[i] < A[i-1] && m[i-1] != 0{
			dp[i] = m[i-1] + 2
		}else if A[i] < A[i-1] && dp[i-1] != 0{
			dp[i] = dp[i-1] + 1
		}else{
			dp[i] = 0
		}

		if dp[i] > max{
			max = dp[i]
		}
	}
	return max
}

func main() {


	a := []int{2,2,2}
	fmt.Println(longestMountain(a))
}
