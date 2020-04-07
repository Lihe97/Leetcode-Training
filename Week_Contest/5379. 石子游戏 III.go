package main

import "fmt"

func stoneGameIII(stoneValue []int) string {
	suma := 0
	sumb := 0
	cur := 0
	flag := 0
	for cur < len(stoneValue){

		if len(stoneValue) - cur >= 3{
			jia,max := help(stoneValue[cur],stoneValue[cur]+stoneValue[cur+1],stoneValue[cur] + stoneValue[cur+1] + stoneValue[cur+2])
			if flag == 0{
				suma += max
				flag = 1
			}else{
				sumb += max
				flag = 0
			}
			cur += jia
		}else if len(stoneValue) - cur == 2{
			if stoneValue[cur+1] > 0{
				max := stoneValue[cur+1] + stoneValue[cur]
				cur += 2
				if flag == 0{
					suma += max
					flag = 1
				}else{
					sumb += max
					flag = 0
				}
			}
		}else if len(stoneValue) - cur == 1{
			max :=  stoneValue[cur]
			if flag == 0{
				suma += max
			}else{
				sumb += max
			}
			break
		}
	}
	if suma > sumb{
		return "Alice"
	}else if suma == sumb{
		return "Tie"
	}else {
		return "Bob"
	}

}
func help(a,b,c int)(int,int){
	if a > b && a > c{
		return 1,a
	}else if b > a && b > c{
		return 2,b
	}else {
		return 3,c
	}
}
func main() {

	a := []int{-1,-2,-3}
	fmt.Println(stoneGameIII(a))

	
}
