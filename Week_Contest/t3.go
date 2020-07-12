package main

import "fmt"



func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {

	flag := make([]bool,len(edges))


	var temp float64
	temp = 1
	var max float64

	dfs(edges,succProb,start,end,flag,temp,&max)
	return  max

}
func dfs(edges [][]int,succ []float64,start,end int,flag []bool,temp float64,max *float64){

	//fmt.Println("?")
	for i := 0 ; i < len(edges) ; i ++{
		if flag[i] == false && (edges[i][0] == start || edges[i][1] == start){
			//fmt.Println(start,temp)
			flag[i] = true
			t := temp
			temp *= succ[i]
			if edges[i][0] == start {
				if edges[i][1] == end{
					if temp > *max{
						*max = temp
					}
				}else{
					dfs(edges,succ,edges[i][1],end,flag,temp,max)
				}
			}else if edges[i][1] == start{
				if edges[i][0] == end{
					if temp > *max{
						*max = temp
					}
				}else{
					dfs(edges,succ,edges[i][0],end,flag,temp,max)
				}
			}
			temp = t
			flag[i] = false
		}
	}

}

func main() {

	fmt.Println(maxProbability(5,[][]int{{2,3},{1,2},{3,4},{1,3},{1,4},{0,1},{2,4},{0,4},{0,2}},[]float64{0.06,0.26,0.49,0.25,0.2,0.64,0.23,0.21,0.77},0,3))
	
}
//5
//[[1,4],[2,4],[0,4],[0,3],[0,2],[2,3]]
//[0.37,0.17,0.93,0.23,0.39,0.04]
//3
//4

//5
//[[2,3],[1,2],[3,4],[1,3],[1,4],[0,1],[2,4],[0,4],[0,2]]
//[0.06,0.26,0.49,0.25,0.2,0.64,0.23,0.21,0.77]
//0
//3
