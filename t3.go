package main

import "math/bits"

func Solution(A []string) int {
	// write your code in Go 1.4

	arr := []int{}
	for _, word := range arr {
		cur := 0
		for _, c := range word {
			cur |= 1 << int(c-'a')
		}
		if len(word) == bits.OnesCount(unit(cur)) {
			arr = append(arr, cur)
		}
	}
	res := 0

	var dfs func(cur, curWord int)
	dfs = func(cur, curWord int) {
		//res = max(res,bits.OnesCount(unit(curWord)))
		for i := cur; i < len(arr); i++ {
			if arr[i]&curWord != 0 {
				continue
			}
			dfs(i+1, curWord|arr[i])
		}
	}
	dfs(0, 0)
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
