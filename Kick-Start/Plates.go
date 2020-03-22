package main

import "fmt"
func pp(arr [][]int,N int ,K int ,P int){
	fmt.Println(arr)
	fmt.Println(N)
	fmt.Println(K)
	fmt.Println(P)

}

func main() {
	var count int
	var N int
	var K int
	var P int
	fmt.Scan(&count)
	//fmt.Println(count)
	//res := []string{}
	for i:= 0 ; i < count ; i ++{
		fmt.Scan(&N)
		fmt.Scan(&K)
		fmt.Scan(&P)
		arr := [][]int{}
		for j := 0 ; j < N ; j++{
			temp := []int{}

			for q := 0 ; q < K ; q ++{
				var tempp int
				fmt.Scan(&tempp)
				temp = append(temp, tempp)
			}
			arr = append(arr, temp)
		}

		pp(arr, N,K,P)
	}

}
