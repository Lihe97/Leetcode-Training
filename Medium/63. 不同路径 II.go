package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	height := len(obstacleGrid)
	width := len(obstacleGrid[0])

	dp := [][]int{}
	for i := 0 ; i < height ; i++{
		temp := make([]int,width)
		dp = append(dp, temp)
	}
	for i := 0 ; i < height ; i++{
		if obstacleGrid[i][0] == 0{
			dp[i][0] = 1
		}else{
			break
		}
	}
	for i := 0 ; i < width ; i++{
		if obstacleGrid[0][i] == 0{
			dp[0][i] = 1
		}else{
			break
		}
	}
	for i := 1 ; i < height ; i ++{

		for j := 1 ; j < width ; j++{
			if obstacleGrid[i][j] == 1{
				dp[i][j] = 0
			}else{
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	//fmt.Println(dp)
	return dp[height-1][width-1]
}
func main() {
	a := [][]int{{0,0,0},{0,1,0},{0,0,0}}
	fmt.Println(uniquePathsWithObstacles(a))
	
}
