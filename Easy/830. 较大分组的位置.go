package main

//func largeGroupPositions(s string) [][]int {
//
//	res := [][]int{}
//	i := 0
//	cnt := 1
//	for i < len(s) {
//		if i == len(s) - 1 || s[i] != s[i+1]{
//			if cnt >=3 {
//				res = append(res,[]int{i - cnt + 1,i})
//			}
//			cnt = 1
//		}else{
//			cnt ++
//		}
//		i ++
//	}
//	return res
//}

func largeGroupPositions(s string) [][]int {


	res := [][]int{}

	start := 0

	for i := 1; i < len(s) ; i ++{
		if i < len(s) && s[i] == s[start]{
			continue
		}
		if i - start + 1 >=3 {
			res = append(res,[]int{start,i})
		}
	}

}

func main() {
	
}
