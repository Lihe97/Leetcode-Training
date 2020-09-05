package main

import "strconv"

func addStrings(num1 string, num2 string) string {


	add := 0
	res := ""

	for i,j := len(num1)-1,len(num2)-1 ; i >= 0 || j >= 0 || add != 0 ; i,j = i-1,j-1{
		x := 0
		y := 0
		if i >= 0 {
			x = int(num1[i]-'0')
		}
		if j >= 0{
			y = int(num2[j] - '0')
		}
		temp := x + y + add
		res = strconv.Itoa(temp%10) + res
		add = temp/10
	}
	return res

}
func main() {
	
}
