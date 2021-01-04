package main

func fib(n int) int {

	num := make([]int,31)
	num[0] = 0
	num[1] = 1
	num[2] = 1

	for i := 3 ; i < len(num) ; i ++{
		num[i] = num[i-1] + num[i-2]
	}
	return num[n]

}


func main() {
	
}
