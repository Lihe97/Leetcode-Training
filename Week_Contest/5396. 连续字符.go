package main
func maxPower(s string) int {

	temp := []byte(s)
	if len(temp) == 1{
		return 1
	}

	res := 0

	cur := 1

	for i := 1 ; i < len(temp) ; i++{
		if temp[i] == temp[i-1]{
			cur ++
		}else{
			cur = 1
		}
		if cur > res{
			res = cur
		}
	}


	return res

}

func main() {
	
}
