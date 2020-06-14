package main

import (
	"fmt"
	"sort"
)

func findLeastNumOfUniqueInts(arr []int, k int) int {

	mp := map[int]int{}
	for i := 0 ; i < len(arr) ; i++{
		mp[arr[i]] ++
	}
	temp := []int{}
	sum := 0
	for _,a := range mp{
		temp = append(temp,a)
		sum ++
	}
	res := 0
	sort.Ints(temp)
	//fmt.Println(temp)
	for   len(temp) > 0 && k >= temp[0]{
		//fmt.Println("???")
		res ++
		k -= temp[0]
		temp = temp[1:]
	}
	//fmt.Println(res)
	return sum - res

}

func main() {

	a := []int{1}
	fmt.Println(findLeastNumOfUniqueInts(a,1))
	
}
