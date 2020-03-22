package main

import (
	"fmt"
	"sort"
)
func qprint(i int,num int, money int ,arr []int)string{

	count := 0
	sort.Ints(arr)
	for i:= 0 ; i < len(arr) ; i ++{
		if money >= arr[i]{
			money -= arr[i]
			count++
		}else {
			break
		}
	}
	res := fmt.Sprintf("Case #%d: %d",i+1,count)
	//fmt.Println(res)
	return res

}
func main() {
	var count int
	var num int
	var money int
	fmt.Scan(&count)
	//fmt.Println(count)
	res := []string{}
	for i:= 0 ; i < count ; i ++{
		fmt.Scan(&num)
		fmt.Scan(&money)
		arr := make([]int,num)
		for j := 0 ; j < num ; j++{
			fmt.Scan(&arr[j])
		}

		res = append(res, qprint(i,num,money,arr))
	}
	for i := 0 ; i < len(res); i++{
		fmt.Println(res[i])
	}
}
