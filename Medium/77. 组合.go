package main

import "fmt"

func combine(n int, k int) [][]int {

	if n <= 0 || k <= 0 || n < k {
		return [][]int{}
	}
	res := [][]int{}



	dfs(n,k,1,[]int{},&res)
	return res
}
func dfs(n,k,cur int, pth []int,res *[][]int){


	if len(pth) == k{
		t := make([]int, k)
		copy(t, pth)
		*res = append(*res, t)
		return
	}
	for i := cur; i <= n; i++ {
		dfs(n,k,i + 1, append(pth, i),res)
	}

}
func main() {

	fmt.Println(combine(4,2))
}
