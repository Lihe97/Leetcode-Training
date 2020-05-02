package main

import (
	"fmt"
	"strconv"
)

func maxDiff(num int) int {

	str  := strconv.Itoa(num)
	bte := []byte(str)
	max,min := false,false
	ma,mi := 0,0
	i := 0
	for i < len(bte) && (!max || !min){
		if max == false && bte[i] != '9'{
			max = true
			ma = i
		}
		if min == false {
			if bte[i] != '1' && i == 0{
				min = true
				mi = i
			}else if bte[i] != '0' && i > 0 && bte[i] != bte[0]{
				min = true
				mi = i
			}
		}
		i ++
	}

	temp1 := make([]byte,len(bte))
	temp2 := make([]byte,len(bte))


	copy(temp1,bte)
	copy(temp2,bte)

	for i := 0 ; i < len(bte) ; i++{
		if temp1[i] == bte[ma]{
			temp1[i] = '9'
		}
		if temp2[i] == bte[mi]{
			if mi == 0{
				temp2[i] = '1'
			}else{
				temp2[i] = '0'
			}
		}
	}

	str1 := string(temp1)
	str2 := string(temp2)
	da, _ := strconv.Atoi(str1)
	xiao , _ := strconv.Atoi(str2)
	fmt.Println(da,xiao)
	return da - xiao


}

func main() {

	fmt.Println(maxDiff(1101057))
	
}
