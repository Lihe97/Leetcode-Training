package main

import "fmt"

func countSquares(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	dp := [][]int{}
	for i := 0 ; i <= len(matrix);i++{
		temp := make([]int,len(matrix[0])+1)
		dp = append(dp,temp)
	}

	res := 0

	for i := 0 ; i < len(matrix); i++{
		for j := 0 ; j < len(matrix[0]) ; j++{
			if matrix[i][j] == 1{

				temp := min(dp[i][j],dp[i+1][j])
				temp = min(temp,dp[i][j+1])
				dp[i+1][j+1] = temp + 1
				res += dp[i+1][j+1]
			}
		}
	}



	return res

}
func min(a,b int)int{
	if a > b{
		return b
	}else{
		return a
	}
}

func main() {

	a := [][]int{{0,1,1,1},{1,1,1,1},{0,1,1,1}}
	fmt.Println(countSquares(a))
	
}
