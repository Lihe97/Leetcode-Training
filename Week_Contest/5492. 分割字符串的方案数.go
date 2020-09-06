package main

import (
	"math"
	"strings"
)

func numWays(s string) int {

	temp := strings.Count(s,"1")
	temp2 := strings.Count(s,"0")




	if temp % 3 != 0{
		return 0
	}
	if temp2 == 0{
		return 1
	}
	if temp == 0 && temp2 != 0{
		sum := 0
		for i := 1 ; i <= len(s) - 2; i ++{
			sum += (len(s)- i - 1)
		}
		return sum%(int(math.Pow10(9))+7)
	}
	t := temp/3
	//fmt.Println(temp,t)

	left := 0
	right := len(s) - 1
	cntl := 0
	cntr := 0
	l1,l2 := 1<<32,0
	for cntl < t + 1{
		//fmt.Println(cntl,left)

		if s[left] == '1'{
			cntl ++
		}
		if cntl == t {
			l1 = min(left,l1)
		}
		if cntl == t + 1{
			l2 = left
		}
		left ++
	}
	r1,r2 := -1<<32,0
	for cntr < t + 1{
		if s[right] == '1'{
			cntr ++
		}
		if cntr == t{
			r1 = max(right,r1)
		}
		if cntr == t{
			r2 = right
		}
		right --
	}
	//fmt.Println(l1,l2,r2,r1)
	return (l2-l1)*(r1-r2+1)%(int(math.Pow10(9))+7)

}
func max(a,b int )int{
	if a > b{
		return a
	}else{
		return b
	}
}
func min(a,b int ) int {
	if a < b {
		return a
	}else{
		return b
	}
}
func main() {
	
}
