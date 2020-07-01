package main

import (
	"fmt"
	"math"
)

func fun(str1,str2 string) int {

	l1,l,r1 := 0,0,0
	max := math.MaxInt16

	for l < len(str2) && l1 < len(str1){
		//fmt.Println(l1,r1,l)

		if str1[r1] == str2[l]{
			l ++
			r1 ++
			if l == len(str2){
				if r1 - l1 < max{
					max = r1 - l1
				}
				l1 ++
				r1 = l1
				l = 0
			}
		}else{
			r1 ++
			if r1 == len(str1){
				l1 ++
				r1 = l1
				l = 0
			}
		}
	}
	if max == math.MaxInt16{
		return 0
	}
	return max
}

func main() {

	var str1 string
	var str2 string
	fmt.Scanln(&str1)
	fmt.Scanln(&str2)

	fmt.Scanln(fun(str1,str2))
	fmt.Println(fun(str1,str2))
	//fmt.Println(fun("12345","344"))

}
