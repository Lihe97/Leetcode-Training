package main

import "strconv"

func getPermutation(n int, k int) string {
	res := []string{}
	flag := make([]bool,n)
	dfs(&res,flag,"",n)
	return res[k-1]
}

func dfs(res *[]string,flag []bool,cur string,n int){
	if len(cur) == n {
		*res = append(*res,cur)
		return
	}
	for i := 0 ; i < len(flag) ; i ++{

		if !flag[i]{

			flag[i] = true
			dfs(res,flag,cur + strconv.Itoa(i+1),n)
			flag[i] = false
		}
	}
}

func main() {
	
}
