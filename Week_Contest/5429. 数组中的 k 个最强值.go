package main

import (
	"fmt"
	"math"
	"sort"
)

//func getStrongest(arr []int, k int) []int {
//
//	var mid int
//	sort.Ints(arr)
//	if len(arr) % 2 == 0{
//		mid = arr[len(arr)/2-1]
//	}else{
//		mid = arr[len(arr)/2]
//	}
//
//
//
//	mp := map[int][]int{}
//
//	temp := []int{}
//
//	for i := 0 ; i < len(arr) ; i ++{
//		mp[int(math.Abs(float64(arr[i]) - float64(mid)))] = append(mp[int(math.Abs(float64(arr[i]) - float64(mid)))],arr[i])
//	}
//	for a := range mp{
//		temp = append(temp,a)
//		sort.Slice(mp[a], func(i, j int) bool {
//			return mp[a][i] > mp[a][j]
//		})
//	}
//	sort.Slice(temp, func(i, j int) bool {
//		return temp[i] > temp[j]
//	})
//	res := []int{}
//
//	for i := 0 ; i < len(temp) ; i++{
//		res = append(res,mp[temp[i]]...)
//	}
//
//	return res[:k]
//}

func getStrongest(arr []int, k int) []int {

	sort.Ints(arr)
	m := arr[(len(arr)-1)/2]

	sort.Slice(arr, func(i, j int) bool {
		a := math.Abs(float64(arr[i])-float64(m))
		b := math.Abs(float64(arr[i])-float64(m))
		if a != b{
			return a > b
		}
		return arr[i] > arr[j]

	})
	return arr[:k]

}

func main() {

	a := []int{1,1,3,5,5}
	fmt.Println(getStrongest(a,2))
	
}
