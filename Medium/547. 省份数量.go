package main


type ud struct {
	count int
	parent []int
}

func constructor(n int) *ud {
	parent := make([]int,n)
	for i := 0 ; i < len(parent); i ++{
		parent[i] = i
	}
	return &ud{
		count:  n,
		parent: parent,
	}
}
func(u *ud) find(x int) int{

	for x != u.parent[x]{
		u.parent[x] = u.find(u.parent[x])
		x = u.parent[x]
	}
	return x
}
func(u *ud) union(x,y int) {
	rootX := u.find(x)
	rootY := u.find(y)
	if rootX == rootY{
		return
	}
	u.parent[rootX] = rootY
	u.count --

}
func findCircleNum(isConnected [][]int) int {
	u := constructor(len(isConnected))

	for i := 0 ; i < len(isConnected) ; i ++{
		for j := i + 1 ; j < len(isConnected) ; j ++{
			if isConnected[i][j] == 1{
				u.union(i,j)
			}
		}
	}
	return u.count
}

func main() {
	
}
