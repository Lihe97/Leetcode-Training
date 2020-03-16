package main

func climbStairs(n int) int {
	return climb(n)
}
func climb(n int) int {
	if n == 1 {
		return 1
	}
	one, two := 1, 2
	for i := 3; i <= n; i++ {
		one, two = two, one + two
	}
	return two
}

func main() {
	
}
