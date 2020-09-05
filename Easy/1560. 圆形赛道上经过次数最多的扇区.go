package main

func mostVisited(n int, rounds []int) []int {

	res := []int{}

	mp := map[int]int{}
	M := 0


	for i := 1 ; i < len(rounds) ; i ++{
		if rounds[i] > rounds[i-1]{
			for j := rounds[i-1] + 1; j <= rounds[i] ; j ++{
				mp[j] ++
				if mp[j] > M{
					M = mp[j]
				}
			}
		}else{
			for j := 1 ; j <= rounds[i] ; j ++{
				mp[j] ++
				if mp[j] > M{
					M = mp[j]
				}
			}
			for j := rounds[i-1] + 1 ; j <= n ; j ++{
				mp[j] ++
				if mp[j] > M{
					M = mp[j]
				}
			}
		}
	}
	mp[rounds[0]] ++
	if mp[rounds[0]] > M{
		M = mp[rounds[0]]
	}
	for b,a := range mp{
		if a == M{
			res = append(res,b)
		}
	}
	sort.Ints(res)
	return res

}


func main() {
	
}
