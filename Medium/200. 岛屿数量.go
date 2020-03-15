package main

import "fmt"
var x  = []int{-1,1,0,0}
var y  = []int{0,0,-1,1}
func numIslands(grid [][]byte) int {

	height := len(grid)
	if height == 0{
		return 0
	}
	width := len(grid[0])
	work := [][]byte{}
	res := 0

	for i:= 0 ; i< height; i++{
		temp := []byte{}
		for j:= 0 ; j < width;j++{
			temp =append(temp, 0)
		}
		work = append(work, temp)
	}
	for i := 0 ; i < height; i++{
		for j:= 0 ; j < width ; j++{
			if grid[i][j] == '1' && work[i][j] == 0{
				DFS(work,grid,i,j)
				res ++
			}
			//fmt.Println(work)
			//fmt.Println(res)
		}

	}
	//fmt.Println(work)
	return res
}
func DFS(work [][]byte,grid [][]byte,i int,j int){
	work[i][j] = 1

	for k := 0 ; k < 4; k ++{

		newx := i + x[k]
		newy := j + y[k]
		if newx < 0 || newx >= len(work) || newy < 0 || newy >= len(work[0]){
			continue
		}
		//fmt.Println(newx,newy)
		//fmt.Println("work:",work[newx][newy])
		//fmt.Println("grid:",grid[newx][newy])
		if work[newx][newy] == 0 && grid[newx][newy] == '1'{
			//fmt.Println(newx,newy)
			//
			//fmt.Println("work[newx][newy]:",work[newx][newy])
			DFS(work,grid,newx,newy)
		}
	}


}
func main() {


	a := [][]byte{{'1','0','1','1','0'},
		          {'1','1','0','1','0'},
		          {'1','0','0','0','0'},
		          {'0','0','0','0','0'}}
	fmt.Println(numIslands(a))
	
}
