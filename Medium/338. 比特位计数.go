package main

//func countBits(num int) []int {
//
//	res := make([]int,num + 1)
//	res[0] = 0
//	for i := 1; i <= num ; i ++{
//		res[i] = res[i & (i-1)] + 1
//	}
//	return res
//
//}
func countBits(num int) []int {

	res := make([]int,num + 1)
	res[0] = 0
	for i := 1; i <= num ; i ++{
		if i % 2 == 1{
			res[i] = res[i-1] + 1
		}else {
			res[i] = res[i/2]
		}
	}
	return res

}
func main() {
	
}
