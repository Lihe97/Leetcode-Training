package main

type UnionFind struct {
	
	parent []int
	weight []float64
}

func newUnionFind(n int) *UnionFind {
	parent := make([]int,n)
	weight := make([]float64,n)

	for i := 0 ; i < n ; i ++{
		parent[i] = i
		weight[i] = 1.0
	}
	return &UnionFind{
		parent: parent,
		weight: weight,
	}
}

func(u *UnionFind) union(x int,y int,val float64){
	rootX := u.find(x)
	rootY := u.find(y)
	if rootX == rootY{
		return
	}
	u.parent[rootX] = rootY
	u.weight[rootX] = u.weight[y] * val / u.weight[x]
}
func(u *UnionFind) find (x int)  int{

	if x != u.parent[x]{
		origin := u.parent[x]
		u.parent[x] = u.find(u.parent[x])
		u.weight[x] *= u.weight[origin]
	}
	return u.parent[x]

}

func(u *UnionFind) isConnected(x int,y int) float64{
	rootX := u.find(x)
	rootY := u.find(y)
	if rootX == rootY{
		return u.weight[x] / u.weight[y]
	}else{
		return -1.0
	}
}


func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {

	unionFind := newUnionFind(2 * len(equations))
	mp := map[string]int{}
	id := 0
	for i := 0 ; i < len(equations) ; i ++{
		a := equations[i][0]
		b := equations[i][1]

		if _,ok := mp[a]; !ok{
			mp[a] = id
			id ++
		}
		if _,ok := mp[b]; !ok{
			mp[b] = id
			id ++
		}
		unionFind.union(mp[a],mp[b],values[i])
	}

	res := []float64{}

	for i := 0 ; i  < len(queries) ; i ++{
		a := queries[i][0]
		b := queries[i][1]
		_,ok1 := mp[a]
		_,ok2 := mp[b]
		if !ok1 || !ok2{
			res = append(res,-1.0)
		}else{
			res = append(res,unionFind.isConnected(mp[a],mp[b]))
		}
	}
	return res
}
func main() {
	
}
