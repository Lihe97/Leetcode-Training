package main

import (
	"fmt"
)

func subsets(nums []int) [][]int {

	res:= make([][]int,1<<len(nums))
	if nums == nil{
		res[0] =[]int{}
		return res
	}
	res[0] = []int{nums[0]}
	//res[1] = []int{}
	for i :=1 ;i <len(nums); i++{
		fmt.Println(res)
		fun(res,1<<i-1,nums[i])
	}
	res[1<<len(nums) -1] = []int{}
	return res
}
func fun(a [][]int,l int,x int){
	for i:=0 ; i <l ;i++{
		//a[i+l] = []int{}
		//fmt.Println(a[i],x)


		a[i+l] = append(a[i], x)
		//fmt.Println(a[i+l])
	}
	a[2*l] =[]int{x}

}

func main() {
	nums:=[]int{1,2,3,4,5,6}

	fmt.Println(subsets(nums))

}
