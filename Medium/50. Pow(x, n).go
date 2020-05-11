package main
//func myPow(x float64, n int) float64 {
//
//	var res float64
//	res = 1
//	if n < 0{
//		x = 1/x
//		n = -n
//	}
//
//	for i := 0 ; i < n ;  i ++{
//		res *= x
//	}
//	return res
//}
func myPow(x float64, n int) float64 {

	if n == 0{
		return 1
	}
	if n < 0 {
		return 1 / myPow(x,-n)
	}
	if n % 2 != 0 {
		return x * myPow(x,n-1)
	}else{
		return myPow(x*x , n/2)
	}


}

func main() {
	
}
