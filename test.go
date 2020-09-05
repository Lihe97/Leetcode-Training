package main

func numOfSubarrays(arr []int) int {

	res := 0
	cnt := 0
	for i := 0 ; i < len(arr) ; i ++{
		if arr[i] % 2 == 1{
			cnt ++
		}
	}
	cnt2 := len(arr) - cnt
	if cnt == 0{
		return 0
	}
	for j := 1 ; j <= cnt ; j += 2{
		res += f(cnt)/(f(j)*f(cnt-j))
	}
	return res

}
func f(x int) int{
	y := 1
	for i := 1 ; i <= x; i ++{
		y = y*i
	}
	return y
}


func main() {

	v


}


