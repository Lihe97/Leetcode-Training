package main

func numJewelsInStones(J string, S string) int {

	res := 0

	mp := map[byte]bool{}

	for i := 0 ; i  < len(J) ; i ++{
		mp[J[i]] = true
	}

	for i := 0 ; i < len(S) ; i ++{
		if mp[S[i]]{
			res ++
		}
	}
	return res

}

func main() {
	
}
