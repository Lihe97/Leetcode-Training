package main

import "fmt"

func longestCommonPrefix(strs []string) string {

	res := strs[0]
	for i := 1 ; i < len(strs) ; i ++{
		res = prefix(res,strs[i])
	}
	return res

}
func prefix(a,b string)string{
	i := 0
	for i < len(a) && i < len(b) && a[i] == b[i]{
		i++
	}
	return a[:i]

}

func main() {
	fmt.Println(prefix("flower","flow"))
	
}
