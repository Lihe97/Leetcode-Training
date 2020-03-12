package main

import (
	"fmt"
	"strconv"
)
func removeKdigits(num string, k int) string {
	if k == 0{
		return num
	}
	if len(num) == 1{
		return ""
	}
	data := []byte(num)

	a := 0
	start , end := 0,1
	for end < len(num){
		if a == k{
			break
		}
		if data[end] > data[start]{
			a++
		}else if data[end] < data[start]{
			start = end
			a++
		}else if data[end] == data[start]{
			start++
		}
		end ++
	}
	fmt.Println("start ",start)
	fmt.Println("a ",a)
	fmt.Println("end ",end)
	res := string(data[start]) + string(data[end:])
	fmt.Println("res",res)
	s,_ := strconv.Atoi(res)
	return  strconv.Itoa(s)
}
func main() {
	s := "5337"
	//src := []rune(s)
	fmt.Println(removeKdigits(s,2))
	
}
