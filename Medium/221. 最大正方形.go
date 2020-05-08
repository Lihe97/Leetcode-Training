package main

import "fmt"

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	dp := [][]int{}
	for i := 0 ; i <= len(matrix) ;i++{

		temp := make([]int,len(matrix[0])+1)
		dp = append(dp,temp)
	}
	max := 0



	for i := 0 ; i < len(matrix) ; i ++{
		for j := 0 ; j < len(matrix[i]) ; j++{
			if matrix[i][j] == '1'{

				temp := min(dp[i+1][j],dp[i][j])
				temp = min(temp,dp[i][j+1])
				dp[i+1][j+1] = temp + 1
				if dp[i+1][j+1]*dp[i+1][j+1] > max{
					max = dp[i+1][j+1] * dp[i+1][j+1]
				}
			}
		}
	}
	//fmt.Println(dp)
	return max

}
func min(a,b int)int{
	if a > b{
		return b
	}else{
		return a
	}
}

func main() {

	b := [][]byte{{'1','0','1','0','0'},{'1','0','1','1','1'},{'1','1','1','1','1'},{'1','0','0','1','0'}}
	//b := [][]byte{{'1'}}
	fmt.Println(maximalSquare(b))
}
