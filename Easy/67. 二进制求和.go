package main

import (
	"fmt"
)

func addBinary(a string, b string) string {

	res := ""
	i,j := len(a)-1,len(b) -1
	flag := 0

	for i >= 0 || j >= 0 {

		if i >= 0 {
			flag += int(a[i]-'0')
			i --
		}
		if j >= 0 {
			flag += int(b[j]-'0')
			j --
		}
		switch flag{
		case 1 :
			res = "1" + res;flag = 0
		case 2:
			res = "0" + res;flag = 1
		case 3:
			res = "1" + res;flag = 1
		case 0:
			res = "0" + res;flag = 0
		}
	}
	if flag == 1{
		res = "1" + res
	}
	return res
}
func main() {

	fmt.Println(addBinary("11","1100"))
	fmt.Println(3<<3)
}
