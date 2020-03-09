package main

import (
	"fmt"
	"math"
)
//暴力+一次遍历

func maxProfit(prices []int) int {
	max := 0
	temp:=0
	for j := len(prices)-1;j>=1 ; j--{

		if prices[j-1] < prices[j]  {
			temp= prices[j] - prices[j-1]
			for i := j+1;i<len(prices);i++{
				if prices[i] -prices[j-1] >temp{
					temp = prices[i] -prices[j-1]
				}
			}
			if temp > max{
				max = temp
			}

		}
	}
	return max
}
func maxProfit1(prices []int) int {
	max := 0
	minPoint := math.MaxInt64
	for i:=0;i<len(prices);i++ {
		if prices[i] < minPoint {
			minPoint = prices[i]
		} else if prices[i] - minPoint > max {
			max = prices[i] - minPoint
		}

	}
	return max
}


func main() {
	a:=[]int{7,1,5,3,6,4}
	fmt.Println(maxProfit(a))
	
}
