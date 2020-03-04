package main

import "fmt"

func strWithout3a3b(A int, B int) string {
	ww := A + B
	ans := ""
	for A > 0 || B > 0 {
		if A > B {
			ans += "aab"
			A -= 2
			B -= 1
		}
		if A < B {
			ans += "bba"
			A -= 1
			B -= 2
		}
		if A == B {
			ans += "ba"
			A -= 1
			B -= 1
		}
	}
	return ans[:ww]

}
func main() {
	fmt.Println(strWithout3a3b(4,1))
	
}
