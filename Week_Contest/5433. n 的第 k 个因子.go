package main

func kthFactor(n int, k int) int {

	res := []int{}

	for i := 1 ; i <= n ; i ++{
		if n%i == 0{
			res = append(res,i)
		}
	}
	if k > len(res){
		return -1
	}
	return res[k-1]

}
func main() {


	
}
