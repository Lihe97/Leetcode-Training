package main

func isValidSudoku(board [][]byte) bool {



	for i := 0 ; i < 9 ; i ++{
		flag := make([]bool,9)
		for j := 0 ; j < 9 ; j ++{
			if board[i][j] != '.'{
				if flag[board[i][j]-'1'] {
					return false
				}else{
					flag[board[i][j]-'1'] = true
				}
			}
		}
	}

	for i := 0 ; i < 9 ; i ++{
		flag := make([]bool,9)
		for j := 0 ; j < 9 ; j ++{
			if board[j][i] != '.'{
				if flag[board[j][i]-'1'] {
					return false
				}else{
					flag[board[j][i]-'1'] = true
				}
			}
		}
	}
	for i := 0 ; i < 9 ; i += 3{
		for j := 0 ; j < 9 ; j += 3{
			flag := make([]bool,9)
			for m := 0 ; m < 3; m ++{
				for n := 0 ; n < 3 ; n ++{
					if board[i+m][j+n] != 'x'{
						if flag[board[i+m][j+n]-'1'] {
							return false
						}else{
							flag[board[i+m][j+n]-'1'] = true
						}
					}
				}
			}
		}
	}
	return true


}

func main() {
	
}
