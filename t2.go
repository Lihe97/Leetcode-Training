package main

import (
	"fmt"
	"sort"
)


/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 实现Sort Merge Join
 * @param setA int整型一维数组
 * @param setB int整型一维数组
 * @return int整型二维数组
 */
func sortMergeJoin( setA []int ,  setB []int ) [][]int {
	// write code here

	res := [][]int{}
	temp := [][3]int{}
	for i := 0 ; i < len(setA) ; i ++{
		temp = append(temp,[3]int{setA[i],i,-1})
	}
	for i := 0 ; i < len(setB) ; i ++{
		temp = append(temp,[3]int{setB[i],i,1})
	}

	sort.Slice(temp, func(i, j int) bool {
		return temp[i][0] < temp[j][0]
	})

	for i := 0 ; i < len(temp) ; i ++{
		for j := i + 1 ; j < len(temp) && temp[j][0] == temp[i][0] ; j ++{
			if temp[j][2] != temp[i][2]{
				res = append(res,[]int{temp[i][1],temp[j][1]})
			}
		}
	}

	return res
}
func main() {

	a := []int{1,2,3,3,4,4,5}
	b := []int{1,2,4,4,7,8}
	fmt.Println(sortMergeJoin(a,b))

}