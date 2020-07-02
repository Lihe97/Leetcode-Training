package main

//func kthSmallest(matrix [][]int, k int) int {
//
//	temp := []int{}
//	for i := 0 ; i < len(matrix) ; i ++{
//		for j := 0 ; j < len(matrix[i]) ; j ++{
//			temp = append(temp,matrix[i][j])
//		}
//	}
//	sort.Ints(temp)
//	return temp[k-1]
//}
func kthSmallest(matrix [][]int, k int) int {

	left := matrix[0][0]
	right := matrix[len(matrix)-1][len(matrix[0])-1]
	for left < right{
		mid := (left + right) / 2

		if check(matrix,mid,k){
			right = mid
		}else{
			left = mid + 1
		}
	}
	return left
}
func check(matrix [][]int,mid int,k int) bool {
	sum := 0
	i := len(matrix)-1
	j := 0
	for i >= 0 && j < len(matrix){
		if matrix[i][j] <= mid{
			sum += i + 1
			j ++
		}else{
			i --
		}
	}
	return sum >= k
}

func main() {
	
}
