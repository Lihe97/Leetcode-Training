package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {

	res := 0
	sort.Ints(g)
	sort.Ints(s)
	j:=0
	for i := 0 ; i < len(g); i++{
		for j <len(s){
			if s[j] >= g[i]{
				res++
				j++
				break
			}else{
				j++
			}
		}
	}
	return res
}
func main() {
	a:= []int{1,2}
	b := []int{1,2,3}
	fmt.Println(findContentChildren(a,b))
	
}
