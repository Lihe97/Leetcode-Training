package main

import "fmt"

func reachNumber(target int) int {
	if target<0{
		target = -target
	}
	i:=0
	sum:= 0
	for sum<target || (sum-target)%2 != 0{

		i++
		sum += i
	}
	return i
}
func main() {
	fmt.Println(reachNumber(3))
	
}
