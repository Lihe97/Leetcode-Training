package main

import "fmt"

func buildArray(target []int, n int) []string {

	length := len(target)
	res := []string{}
	i := 1
	j := 0

	for i <= target[length-1]{
		if i == target[j]{
			res = append(res,"Push")
			i++
			j++
		}else{
			res = append(res,"Push")
			res = append(res,"Pop")
			i++
		}
	}


	return res
}

func main() {

	a := []int{2,3,4}
	fmt.Println(buildArray(a,4))
	
}
