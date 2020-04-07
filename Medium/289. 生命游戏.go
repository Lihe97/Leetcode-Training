package main

func gameOfLife(board [][]int)  {
	height := len(board)
	width := len(board[0])
	res := [][]int{}
	dx := []int{-1,-1,-1,0,1,1,1,0}
	dy := []int{-1,0,1,1,1,0,-1,-1}
	for i := 0 ; i < height ; i ++{
		temp := []int{}

		for j := 0 ; j < width ; j++{
			yi := 0
			for  k := 0 ; k < 8 ; k ++{
				newx := i + dx[k]
				newy := j + dy[k]
				if newx >= 0 && newx < height && newy >= 0 && newy < width{
					if board[newx][newy] == 1{
						yi ++
					}
				}
			}
			//fmt.Println(board[i][j],yi)
			if board[i][j] == 0 && yi == 3{
				temp = append(temp, 1)
			}else if board[i][j] == 1 && yi < 2{
				temp = append(temp,0)
			}else if board[i][j] == 1 && yi <= 3{
				temp = append(temp,1)
			}else if board[i][j] == 1 && yi > 3{
				temp = append(temp,0)
			}else{
				temp = append(temp,board[i][j])
			}
		}
		//fmt.Println("temp",temp)
		res = append(res, temp)
	}
	copy(board,res)
	//fmt.Println(board)

}
func main() {

	a := [][]int{{0,1,0},{0,0,1},{1,1,1,},{0,0,0}}
	gameOfLife(a)
	
}
