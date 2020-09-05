package main

import "strconv"


func thousandSeparator(n int) string {
	if n < 1000{
		return strconv.Itoa(n)
	}
	t := strconv.Itoa(n)
	cnt := 0
	res := ""
	for i := len(t) - 1 ; i >= 0 ; i --{
		cnt ++
		if cnt == 3 && i != 0{

			res += string(t[i])
			res += "."
			cnt = 0
		}else{
			res += string(t[i])
		}

	}
	r := ""
	for i := len(res)-1 ; i >=0 ; i --{
		r += string(res[i])
	}
	return r

}

func main() {
	
}
