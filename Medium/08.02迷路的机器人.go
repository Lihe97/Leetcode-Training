package main
var result [][]int //结果
var visted [][]bool //剪枝，因为DFS，如果访问过则说明路不同
var rows int
var cols int
func pathWithObstacles(obstacleGrid [][]int) [][]int {

	rows = len(obstacleGrid)
	if rows == 0 {
		return nil
	}
	cols = len(obstacleGrid[0])
	if cols == 0{
		return nil
	}
	if obstacleGrid[0][0] == 1 || obstacleGrid[rows -1][cols-1] == 1{
		return nil
	}
	result = [][]int{}
	visted = [][]bool{}
	for i:= 0; i < rows ;i++{
		visted = append(visted, make([]bool,cols))
	}
	if path(obstacleGrid, 0, 0){
		return result
	}
	return nil
}

func path(obstacleGrid [][]int, row, col int) bool{
	if row == rows || col == cols{ //越界检查
		return false
	}
	if visted[row][col] || obstacleGrid[row][col] == 1{ //如果访问过或者不同返回
		return false
	}
	visted[row][col] = true //标记当前访问过
	result = append(result, []int{row, col}) //记录结果
	if (row == (rows -1)) && (col == (cols -1)){ //得到结果结束
		return true
	}
	if path(obstacleGrid, row+1,col){ //左树
		return true
	}
	if path(obstacleGrid, row, col+1){ //右树
		return true
	}

	result = result[:len(result)-1] //在结果中移除
	return false
}

