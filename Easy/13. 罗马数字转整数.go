package main

import "fmt"

func romanToInt(s string) int {

	res := 0
	i := 0
	for i < len(s) {
		if s[i] == 'M'{
			res += 1000
			i ++
		}else if s[i] == 'I'{
			if i == len(s) -1{
				res ++
				i ++
			}else if s[i+1] == 'V'{
				res += 4
				i += 2
			}else if s[i+1] == 'X'{
				res += 9
				i += 2
			}else{
				i ++
				res ++
			}
		}else if s[i] == 'X'{
			if i == len(s) -1{
				res += 10
				i ++
			}else if s[i+1] == 'L'{
				res += 40
				i += 2
			}else if s[i+1] == 'C'{
				res += 90
				i += 2
			}else{
				i ++
				res += 10
			}
		}else if s[i] == 'C'{
			if i == len(s) -1{
				res += 100
				i ++
			}else if s[i+1] == 'D'{
				res += 400
				i += 2
			}else if s[i+1] == 'M'{
				res += 900
				i += 2
			}else{
				i ++
				res += 100
			}
		}else if s[i] == 'D'{
			res += 500
			i ++
		}else if s[i] == 'V'{
			res += 5
			i ++
		}else if s[i] == 'L'{
			res += 50
			i ++
		}
	}
	return  res
}

func main() {

	fmt.Println(romanToInt("MDLXX"))
	
}
