package main

import "fmt"

//func findCircleNum(M [][]int) int {
//
//	num := len(M)
//	res := [][]int{}
//	flag := make([]bool,num)
//	for i := 0 ; i < num ; i ++{
//		temp := []int{}
//		if flag[i] == false{
//			DDFS(i,num,&temp,M,flag)
//			fmt.Println(temp)
//			res = append(res, temp)
//		}
//	}
//	//fmt.Println(flag)
//	return len(res)
//}
//
//func DDFS(i int,num int,temp *[]int,M [][]int,flag []bool){
//	*temp = append(*temp, i)
//	flag[i] = true
//	for j := 0 ; j < num; j ++{
//		if M[i][j] == 1 && flag[j] == false{
//			DDFS(j,num,temp,M,flag)
//		}
//	}
//
//}

type UnionFind struct {
	count int
	parent []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int,n)

	for i := 0 ; i < n ; i ++{
		parent[i] = i
	}

	return &UnionFind{
		count:  n,
		parent: parent,
	}
}

func(u *UnionFind) Union(i,j int){
	rootp := u.Find(i)
	rootq := u.Find(j)

	if rootp == rootq{
		return
	}
	u.parent[rootq] = rootp

	u.count --

}
func (u *UnionFind) Find(i int)  int {
	//fmt.Println("before",i)

	for u.parent[i] != i{
		u.parent[i] = u.parent[u.parent[i]]
		i = u.parent[i]
	}
	//fmt.Println("after",i)

	return i

}
func findCircleNum(M [][]int) int {

	if len(M) == 0{
		return 0
	}
	ufind := NewUnionFind(len(M))
	for i := 0 ; i < len(M) ; i ++{
		for j := i + 1 ; j < len(M) ;  j ++{
			if M[i][j] == 1{
				ufind.Union(i,j)
			}
		}
	}
	return ufind.count
}

func main() {

	a := [][]int{{1,0,0,1},{0,1,1,0},{0,1,1,1},{1,0,1,1}}
	fmt.Println(findCircleNum(a))
	
}
