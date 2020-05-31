package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func hasAllCodes(s string, k int) bool {


	for i := 0 ; i < int(math.Pow(2,float64(k))) ; i++{
		q := convertToBin(i)
		if len(q) == k{
			if !strings.Contains(s,q){
				return false
			}
		}else{
			p := ""
			w := k - len(q)
			for i := 0 ; i < w ; i++{
				p += "0"
			}
			q = p + q
			if !strings.Contains(s,q){
				return false
			}
		}

	}

	return true
}
func convertToBin(num int) string {
	s := ""

	if num == 0 {
		return "0"
	}
	for ;num > 0 ; num /= 2 {
		lsb := num % 2
		s = strconv.Itoa(lsb) + s
	}
	return s

}

func main() {


	fmt.Println(hasAllCodes("00110110",2))
	
}
