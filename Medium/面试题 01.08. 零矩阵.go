package main
func setZeroes(matrix [][]int)  {
	//del_h, del_l := make(map[int]bool),  make(map[int]bool)
	h,l := make(map[int]bool),make(map[int]bool)
	for i:= 0; i <len(matrix);i++{
		for j:=0;j<len(matrix[0]);j++{
			if matrix[i][j] == 0{
				h[i] = true
				l[j] = true
			}
		}
	}
	for k,_ := range h {
		for j := 0; j < len(matrix[0]); j ++ {
			matrix[k][j] = 0
		}
	}
	for k, _ := range l {
		for j := 0; j < len(matrix); j ++ {
			matrix[j][k] = 0
		}
	}
}
func main() {
	
}
