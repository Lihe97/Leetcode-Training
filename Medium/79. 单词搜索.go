package main

import "fmt"

//func exist(board [][]byte, word string) bool {
//
//	flag := [][]bool{}
//
//	for i := 0 ; i < len(board) ; i ++{
//		temp := make([]bool,len(board[0]))
//		flag = append(flag,temp)
//	}
//
//	for i := 0 ; i < len(board) ; i ++{
//		for j := 0 ; j < len(board[0]) ; j ++{
//			if board[i][j] == word[0]{
//				flag[i][j] = true
//				if dfs(board,flag,i,j,word[1:]){
//					return true
//				}
//				flag[i][j] = false
//			}
//		}
//	}
//	return false
//}
//func dfs(board [][]byte,flag [][]bool,x,y int,word string) bool{
//
//	if word == ""{
//		return true
//	}
//	dx := []int{-1,1,0,0}
//	dy := []int{0,0,1,-1}
//	for i := 0 ; i < 4 ; i ++{
//		nx := x + dx[i]
//		ny := y + dy[i]
//
//		if nx >= 0 && nx < len(board) && ny >= 0 && ny < len(board[0]) && !flag[nx][ny] && board[nx][ny] == word[0]{
//			flag[nx][ny] = true
//			if dfs(board,flag,nx,ny,word[1:]){
//				return true
//			}
//			flag[nx][ny] = false
//		}
//	}
//	return false
//}

func exist(board [][]byte, word string) bool {

	flag := [][]bool{}

	for i := 0 ; i < len(board) ; i ++{
		temp := make([]bool,len(board[i]))
		flag = append(flag,temp)
	}
	fmt.Print(flag)

	for i := 0 ; i < len(board) ; i ++{
		for j := 0 ;  j < len(board[i]) ; j ++{
			if board[i][j] == word[0]{
				flag[i][j] = true
				if dfs(i,j,word[1:],flag,board){
					return  true
				}
				flag[i][j] = false
			}
		}
	}

	return false
}
func dfs(curx ,cury int,word string,flag [][]bool,board [][]byte) bool {
	if word == ""{
		return true
	}
	dx := []int{-1,1,0,0}
	dy := []int{0,0,1,-1}

	for i := 0  ; i < 4; i ++{
		//newx := curx + dy[i]
		//newy := cury + dx[i]
		//if newx >= 0 && newx < len(flag) && newy >= 0 && newy < len(flag[i]) && !flag[newx][newy]{
		nx := curx + dx[i]
		ny := cury + dy[i]

		if nx >= 0 && nx < len(board) && ny >= 0 && ny < len(board[0]) && !flag[nx][ny] && board[nx][ny] == word[0]{
			flag[nx][ny] = true
			if dfs(nx,ny,word[1:],flag,board){
				return true
			}
			flag[nx][ny] = false
		}
	}
	return false
}
func main() {
	a := [][]byte{{'A','B','C','E'},{'S','F','C','S'},{'A','D','F','E'}}
	fmt.Println(exist(a,"ABCCED"))
}
