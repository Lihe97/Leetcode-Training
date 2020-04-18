package main

import "fmt"

func numWays(n int, relation [][]int, k int) int {

	flag := false
	mp := map[int][]int{}
	for i := 0 ; i < len(relation) ; i ++{
		if relation[i][1] == n-1{
			flag = true
			break
		}
	}
	if flag == false{
		return 0
	}
	for i := 0 ; i < len(relation) ; i ++{
		mp[relation[i][0]] = append(mp[relation[i][0]], relation[i][1])
	}
	res := 0

	DFS(0,mp,k,&res,0,n-1)

	return res
}
func DFS(cur int,mp map[int][]int,k int ,res *int,num int,fin int){
	//fmt.Println(cur,num)
	if cur == fin && num == k{

		*res ++

		return
	}
	if num > k {
		return
	}
	for a,b := range mp{
		if a == cur {

			for i := 0 ; i < len(b) ; i++{
				DFS(b[i],mp,k,res,num+1,fin)
			}
		}
	}


}

func main() {

	//a := [][]int{{0,2},{0,4},{1,4},{2,0},{2,1},{2,3},{3,4}}
	a := [][]int{{0,2},{2,1}}
	fmt.Println(numWays(3,a,2))
	
}
