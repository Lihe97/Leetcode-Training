package main

func integerReplacement(n int) int {


	mp := map[int]bool{}

	queue := []int{}

	queue = append(queue, n)

	mp[n] = true

	cnt := 0

	for len(queue) != 0 {
		tmp := []int{}
		for len(queue) != 0{
			a := queue[0]
			queue = queue[1:]
			if a  == 1{
				return cnt
			}
			if a % 2 == 0{
				if !mp[a/2]{
					mp[a/2] = true
					tmp = append(tmp,a/2)
				}
			}else{
				if !mp[a+1]{
					mp[a+1] = true
					tmp = append(tmp,a+1)
				}
				if !mp[a-1]{
					mp[a-1] = true
					tmp = append(tmp,a-1)
				}
			}
		}
		cnt ++
		queue = tmp
	}
	return cnt
}
