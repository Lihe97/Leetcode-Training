package main

func findSmallestSetOfVertices(n int, edges [][]int) []int {

	mp := map[int]int{}

	for i := 0 ; i < len(edges) ; i ++{
		mp[edges[i][1]] ++
	}
	res := []int{}
	for i := 0 ; i < n ; i ++{
		if mp[i] == 0{
			res = append(res,i)
		}
	}
	return res

}

func main() {
	
}
