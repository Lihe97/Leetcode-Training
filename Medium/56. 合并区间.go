package main

import (
	"fmt"
	"sort"
)
func merge(intervals [][]int) [][]int {

	if len(intervals) ==0 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	res = append(res, intervals[0])

	for i := 1;  i < len(intervals); i ++{
		pre := res[len(res)-1]
		cur := intervals[i]
		if cur[0] > pre[1]{
			res = append(res, cur)
			continue
		}
		if cur[1] > pre[1]{
			pre[1] = cur[1]
		}
	}

	return res
}
//func merge(intervals [][]int) [][]int {
//
//	if len(intervals) == 0{
//		return intervals
//	}
//	res := [][]int{}
//	flag := make([]bool,len(intervals))
//	DDDFS(flag,intervals[0][0],intervals[0][1],intervals,&res)
//	return res
//}
//func DDDFS(flag []bool,left,right int,intervals [][]int,res *[][]int){
//	temp := false
//	for i := 0 ; i < len(intervals) ;  i ++{
//		if flag[i] == false && intervals[i][0] <= right && intervals[i][1]>= left{
//			flag[i] = true
//			left = min(left,intervals[i][0])
//			right = max (right,intervals[i][1])
//			temp = true
//			DDDFS(flag,left,right,intervals,res)
//		}
//	}
//	if temp == false{
//		*res = append(*res, []int{left,right})
//	}
//	q := false
//	for i:= 0 ; i < len(intervals) ; i ++{
//		if flag[i] == false{
//			flag[i] = true
//			left = intervals[i][0]
//			right = intervals[i][1]
//			q = true
//			break
//		}
//	}
//	if q{
//		DDDFS(flag,left,right,intervals,res)
//	}
//
//}
func max(a,b int)int{
	if a > b{
		return a
	}else{
		return b
	}
}


func main() {

	a := [][]int{{1,3},{2,6},{8,10},{15,18}}
	fmt.Println(merge(a))
	
}
