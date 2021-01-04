package main
func numberOfMatches(n int) int {

	res := 0

	for n != 1{
		if n % 2 == 1{


			res += (n-1)/2
			n = (n-1)/2 + 1
		}else{
			n = n/2
			res += n
		}

	}
	return res
}

func main() {
	
}
