package main

import "fmt"

func convert(s string, numRows int) string {

	res := ""
	if numRows == 1{
		return s
	}
	for i := 0 ; i < numRows ; i ++{

		if i == 0 || i == numRows -1{
			for j := i ; j < len(s) ; j += 2*numRows-2{
				res += string(s[j])
			}
		}else{
			for m,n := i, 2*numRows-2-i;m < len(s) || n < len(s);{
				if m <len(s){
					res += string(s[m])
				}
				if n < len(s){
					res += string(s[n])
				}
				m += 2*numRows-2
				n += 2*numRows-2
			}
		}
	}
	return res
}

func main() {

	fmt.Println(convert("LEETCODEISHIRING",3))
	//LCIRETOESIIGEDHN
}
