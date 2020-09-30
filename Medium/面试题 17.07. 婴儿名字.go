package main

import (
	"strconv"
	"strings"
)

type Union struct {
	parent []int
	num []int
	names []string
}

func Constructor(names []string) *Union {
	parent := make([]int,len(names))
	num := make([]int,len(names))

	for i := 0 ; i < len(names);  i ++{
		a := strings.Index(names[i],"(")
		b := strings.Index(names[i],")")
		temp := names[i][a+1:b]
		t, _ := strconv.Atoi(temp)
		num[i] = t
		parent[i] = i
	}
	return &Union{
		parent: parent,
		num:    num,
		names:names,
	}
}
func (u *Union)find(x int) int  {

	for u.parent[x] != x{
		u.parent[x] = u.parent[u.parent[x]]
		x = u.parent[x]
	}
	return x
}
func (u *Union) join (a,b int){
	rootp := u.find(a)
	rootq := u.find(b)
	if rootp == rootq{
		return
	}
	if u.names[rootp] < u.names[rootq]{
		u.parent[rootq] = rootp
		u.num[rootp] += u.num[rootq]
	}else{
		u.parent[rootp] = rootq
		u.num[rootq] += u.num[rootp]
	}
}


func trulyMostPopular(names []string, synonyms []string) []string {

	res := []string{}

	mp := map[string]int{}

	for i := 0 ; i < len(names); i ++{
		temp := strings.Index(names[i],"(")

		mp[names[i][:temp]] = i
	}

	union := Constructor(names)

	for i := 0 ; i < len(synonyms) ; i ++{
		a := strings.Index(synonyms[i],"(")
		b := strings.Index(synonyms[i],",")
		c := strings.Index(synonyms[i],")")
		temp1 := synonyms[i][a+1:b]
		temp2 := synonyms[i][b+1:c]

		union.join(mp[temp1],mp[temp2])
	}
	for i := 0 ; i < len(union.parent) ; i ++{
		if union.parent[i] == i{
			t := strings.Index(names[i],"(")
			temp := names[i][:t]
			temp += "("
			temp += strconv.Itoa(union.num[i])
			temp += ")"
			res = append(res,temp)
		}
	}
	return res

}

func main() {


}
