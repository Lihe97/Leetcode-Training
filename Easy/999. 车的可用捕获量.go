package main

func numRookCaptures(board [][]byte) int {

	res := 0
	dx := []int{1,0,-1,0}
	dy := []int{0,1,0,-1}

	var curx,cury int
	for i := 0 ; i < len(board) ; i ++{
		for j := 0 ; j < len(board[0]) ; j++{
			if board[i][j] == 'R'{
				curx = i
				cury = j
			}
		}
	}

	for i := 0 ; i < 4; i ++{
		for step := 1 ; ; step++{
			newx := curx + step*dx[i]
			newy := cury + step*dy[i]
			if newx < 0 || newx >= 8 || newy < 0 || newy >= 8 || board[newx][newy] == 'R'{
				break
			}
			if board[newx][newy] == 'p'{
				res ++
			}
		}
	}
	return res
}
func main() {
	
}
