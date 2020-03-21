package main

import "fmt"

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {

	res := 0
	for i:= 0 ; i <len(arr1) ; i++{
		temp := true
		for j:= 0 ; j < len(arr2); j++{
			if arr1[i] - arr2[j] <= d && arr2[j] - arr1[i]<=d{
				temp = false
				break
			}
		}
		if temp{
			res ++
		}
	}
	return res

}
func main() {
	arr1:= []int{1,4,2,3}
	arr2 := []int{-4,-3,6,10,20,30}
	fmt.Println(findTheDistanceValue(arr1,arr2,3))
	
}
