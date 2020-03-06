package main
func findContinuousSequence(target int) [][]int {

	res :=[][]int{}

	left := 1

	right := 2
	for left < right {

		num := (left + right)*(right - left + 1)/2
		if num == target{
			tmp := []int{}
			for i:=left ; i<=right ; i++{
				tmp = append(tmp, i)
			}
			res = append(res,tmp)
			left++
		}else {

			if num < target {
				right++
			} else {
				left++
			}
		}


	}
	return res

}
func main() {
	a:= [][]int{}
	a[0] = []int{1}
}
