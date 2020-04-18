package main

import "fmt"

func getTriggerTime(increase [][]int, requirements [][]int) []int {

	C,R,H := 0,0,0
	res := []int{}
	for i := 0 ; i < len(requirements) ; i++{
		res = append(res, -1)
	}


	for i := 0 ; i < len(increase) ; i ++{
		for j := 0 ; j < len(requirements) ; j ++{
			if res[j] == -1 && requirements[j][0] <= C && requirements[j][1] <= R && requirements[j][2] <= H{
				res[j] = i
			}
		}

		C += increase[i][0]
		R += increase[i][1]
		H += increase[i][2]


	}
	for j := 0 ; j < len(requirements) ; j ++{
		if res[j] == -1 && requirements[j][0] <= C && requirements[j][1] <= R && requirements[j][2] <= H{
			res[j] = len(increase)
		}
	}

	return res

}

func main() {


	a := [][]int{{2,8,4},{2,5,0},{10,9,8}}
	b := [][]int{{0,0,0}}
	fmt.Println(getTriggerTime(a,b))
}
