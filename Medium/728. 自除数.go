package main

import "fmt"

func selfDividingNumbers(left int, right int) []int {
	res := []int{}
	for i := left; i<=right; i++{
		if isZichu(i){
			res = append(res, i)
		}
		//fmt.Println(res)
	}
	return res
}

func isZichu(a int)bool{
	c := a
	for c!=0{
		if c%10 == 0{
			return false
		}
		if a%(c%10) !=0 {
			return false
		}
		c = c/10
	}
	return true
}
func main() {
	fmt.Println(selfDividingNumbers(1,22))
	
}
