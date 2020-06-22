package main
func xorOperation(n int, start int) int {

	//res := []int{}
	temp := start
	for i := 1 ; i < n ; i ++{
		//res = append(res,start + 2*i)
		temp ^= start + 2 * i
	}
	return temp

}

func main() {
	
}
