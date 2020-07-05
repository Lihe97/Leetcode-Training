package main

import "fmt"

func numSubmat(mat [][]int) int {


	dp2 := [][]int{}

	for i := 0 ; i < len(mat) ; i ++{

		temp2 := make([]int,len(mat[0]))

		dp2 = append(dp2,temp2)

	}
	res := 0



	for i := 0 ; i < len(mat) ; i ++{
		if mat[i][0] == 1{
			dp2[i][0] = 1
		}
	}
	for i := 1; i < len(mat[0]); i ++{
		for j := 0 ; j < len(mat) ; j ++{
			if mat[j][i] == 1{
				if dp2[j][i-1] != 0{
					dp2[j][i] = dp2[j][i-1] + 1
				}else{
					dp2[j][i] = 1
				}
			}
		}
	}

	for i := 0 ; i < len(mat[0]); i ++{
		if dp2[0][i] != 0{

			res += dp2[0][i]
		}
	}
	fmt.Println(res)
	for i := 1 ; i < len(mat) ; i ++{
		for j := 0 ; j < len(mat[0]) ; j ++{
			if mat[i][j] == 1{
				res += dp2[i][j]
				temp := dp2[i][j]
				for k := i-1 ; k >= 0 ; k --{
					if dp2[k][j] == 0{
						break
					}

					temp = min(temp,dp2[k][j])
					res += temp
					//fmt.Println(res,i,j)
				}
			}
		}
	}
	fmt.Println(dp2)


	return res
}
func min(a,b int)int{
	if a < b{
		return a
	}else{
		return b
	}
}
func main() {

	//a := [][]int{{1,0,1},{0,1,0},{1,0,1}}
	a := [][]int{{1,1,1,1,0,1,0},{1,1,1,0,0,0,1},{0,1,1,1,1,0,0},{1,1,0,1,1,0,1},{1,0,0,0,0,0,1},{1,1,0,1,1,1,1},{1,1,0,0,1,1,1}}
	//a := [][]int{{0,1,1,0},{0,1,1,1},{1,1,1,0}}
	fmt.Println(numSubmat(a))
	
}
