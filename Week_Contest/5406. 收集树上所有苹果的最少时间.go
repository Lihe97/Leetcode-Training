package main

import "fmt"

func minTime(n int, edges [][]int, hasApple []bool) int {

	res := 0
	flag := make([]bool,n)
	flag[0] = true

	mp := map[int]int{}

	for i := 0 ; i < len(edges) ; i ++{
		mp[edges[i][1]] = edges[i][0]
	}


	for i :=  0 ; i < len(hasApple) ; i ++{
		if hasApple[i] {
			dfsedge(i,&res,flag,mp)
		}
	}
	return res * 2
}
func dfsedge(to int , res *int,flag []bool,mp map[int]int){

	if flag[to] == false{
		*res ++
		flag[to] = true
		dfsedge(mp[to],res,flag,mp)
	}
}

func main() {

	edges := [][]int{{0,1},{0,2},{1,4},{1,5},{2,3},{2,6}}
	hasApple := []bool{false,false,true,false,true,true,false}
	fmt.Println(minTime(7,edges,hasApple))

}
