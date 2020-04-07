package main

import "fmt"

func numSteps(s string) int {

	var sum int64
	sum = 0
	res := 0
	for i:= 0 ; i < len(s);i++{
		sum += (int64(s[i])-48) << (len(s)-i-1)

	}
	for sum != 1{
		if sum % 2 ==1{
			sum ++
		}else{
			sum /= 2
		}
		res ++
	}
	return res

}

func main() {

	fmt.Println(numSteps("1111110011101010110011100100101110010100101110111010111110110010"))
}
