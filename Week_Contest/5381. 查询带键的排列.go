package main

import "fmt"

func processQueries(queries []int, m int) []int {

	temp := []int{}
	for i := 1 ; i <= m ; i ++{
		temp = append(temp, i)
	}
	res := []int{}

	for i := 0 ; i < len(queries) ; i ++{
		for j := 0 ; j < len(temp) ; j++{
			if temp[j] == queries[i]{
				res = append(res, j)
				q := []int{}
				q = append(q,temp[j])
				q = append(q,temp[0:j]...)
				q = append(q,temp[j+1:]...)
				temp = q
			}
		}
	}
	return res

}
func main() {
	a := []int{7,5,5,8,3}
	fmt.Println(processQueries(a,8))

	
}
