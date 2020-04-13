package main

import "fmt"

func findCircleNum(M [][]int) int {

	num := len(M)
	res := [][]int{}
	flag := make([]bool,num)
	for i := 0 ; i < num ; i ++{
		temp := []int{}
		if flag[i] == false{
			DDFS(i,num,&temp,M,flag)
			fmt.Println(temp)
			res = append(res, temp)
		}
	}
	//fmt.Println(flag)
	return len(res)
}

func DDFS(i int,num int,temp *[]int,M [][]int,flag []bool){
	*temp = append(*temp, i)
	flag[i] = true
	for j := 0 ; j < num; j ++{
		if M[i][j] == 1 && flag[j] == false{
			DDFS(j,num,temp,M,flag)
		}
	}

}
func main() {

	a := [][]int{{1,0,0,1},{0,1,1,0},{0,1,1,1},{1,0,1,1}}
	fmt.Println(findCircleNum(a))
	
}
