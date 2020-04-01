package main

import "fmt"

func isValid(s string) bool {
	mp := map[byte]byte{
		')' : '(',
		'}' : '{',
		']' : '[',
	}
	res := []byte{}



	for i := 0 ; i < len(s) ; i ++ {

		if len(res) > 0 {
			if res[len(res)-1 ] == mp[s[i]] {
				res = res[:len(res)-1 ]
				continue
			}
		}

		res = append(res, s[i])
	}


	return len(res) == 0
}
func main() {
	fmt.Println(isValid("{}"))
	
}
