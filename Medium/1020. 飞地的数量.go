package main

var dir = [][]int{{0,1},{0,-1},{1,0},{-1,0}}
func numEnclaves(grid [][]int) int {
	m,n := len(grid), len(grid[0])
	flag := make([][]bool,m)
	for i := range flag {
		flag[i] = make([]bool, n)
	}


	for i := 0 ; i < m ; i ++{
		dfsFly(i,0,grid,flag)
		dfsFly(i,n-1,grid,flag)
	}
	for i := 0 ; i < n ; i ++{
		dfsFly(0,i,grid,flag)
		dfsFly(m-1,i,grid,flag)
	}
	res := 0

	for i := 1 ; i < m - 1 ; i++{
		for j := 1 ; j < n - 1 ; j ++{
			if !flag[i][j] && grid[i][j] == 1{
				res ++
			}
		}
	}
	return res
}
func dfsFly(x,y int,grid [][]int,flag [][]bool)  {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == 0 || flag[x][y] {
		return
	}
	flag[x][y] = true
	for _,v := range dir{
		dfsFly(x+v[0],y+v[1],grid,flag)
	}
}
