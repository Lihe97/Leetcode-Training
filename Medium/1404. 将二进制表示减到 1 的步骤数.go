package main

import (
	"fmt"
)

func numSteps(s string) int {

	res := 0
	b := []byte(s)

	for len(b) != 1{
		if b[len(b)-1] == '0'{
			b = b[:len(b) - 1]
			res ++
		}else if b[len(b)-1] == '1'{
			i := len(b)-1
			for  ; i >= 0 ; i --{
				if b[i] == '1'{
					b[i] = '0'
				}else{
					b[i] = '1'
					break
				}
			}
			if i == -1{
				b = append([]byte{'0'},b...)
			}
			res ++
		}

	}
	return res

}

func main() {

	//s := "abc"
	//b := []byte(s)
	fmt.Println(numSteps("1101"))

}
