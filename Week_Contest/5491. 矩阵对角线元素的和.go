package main
func diagonalSum(mat [][]int) int {

	sum := 0
	for i := 0 ; i < len(mat) ; i ++{
		sum += mat[i][i]
	}
	for i := len(mat)-1; i >= 0 ; i-- {
		sum += mat[len(mat)-i-1][i]
	}
	if len(mat) % 2 == 0{
		return sum
	}else{
		sum -= mat[len(mat)/2][len(mat)/2]
		return sum
	}
}

func main() {
	
}
