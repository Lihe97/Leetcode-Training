package main

import "fmt"

//func generateParenthesis(n int) []string {
//	res := []string{}
//	a := n-1
//	b := n
//	temp := 1
//	r := "("
//	ccc(a,b,temp,n,r,&res)
//
//	return res
//}
//func ccc(a,b,temp,n int,res string,result *[]string){
//	if a > 0 || b > 0{
//		if temp == 0{
//			res += "("
//			ccc(a-1,b,1,n,res,result)
//		}else if temp == n{
//			res += ")"
//			ccc(a,b-1,temp-1,n,res,result)
//		}else{
//			q := res
//			if a > 0{
//				res += "("
//				ccc(a-1,b,temp+1,n,res,result)
//			}
//			q += ")"
//			ccc(a,b-1,temp-1,n,q,result)
//		}
//	}else{
//		*result = append(*result, res)
//	}
//}
func generateParenthesis(n int) []string {

	res := []string{}

	dfs(n,0,0,"",&res)

	return res
}
func dfs(n,lc,rc int,temp string,res *[]string){
	if lc == n && rc == n{
		*res = append(*res,temp)
		return
	}else{
		if lc < n{
			dfs(n,lc+1,rc,temp + "(",res)
		}
		if rc < n && lc > rc{
			dfs(n,lc,rc+1,temp+")",res)
		}
	}

}

func main() {

	fmt.Println(generateParenthesis(4))
	
}
