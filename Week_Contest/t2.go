package main

import (
	"fmt"
)

func canArrange(arr []int, k int) bool {

	mp := map[int]int{}
	for i := 0 ; i < len(arr) ; i ++ {
		mp[arr[i]%k] += 1
	}

	//fmt.Println(mp)
	for a,b := range mp{
		//fmt.Println(a,b)
		if a == 0 || b == 0{
			continue
		}
		if mp[-a] > 0 {
			if mp[-a] >= b{
				mp[-a] -= b
				mp[a] = 0
				//fmt.Println("??",b)
			}else if mp[-a] < b && mp[-a] > 0{
				mp[a] -= mp[-a]
				mp[-a] = 0
				//fmt.Println(b)
			}
		}
		//fmt.Println(mp)
		if mp[a] < 0{
			if mp[-k+a%k] >= mp[a]{
				mp[-k+a%k] -= mp[a]
				mp[a] = 0
			}else if mp[-k+a%k] < mp[a] && mp[-k+a%k] > 0{
				mp[a] -= mp[-k+a%k]
				mp[-k+a%k] = 0
			}
		}else  if mp[a] > 0{

			if mp[k-a%k] >= mp[a]{
				mp[k-a%k] -= mp[a]
				mp[a] = 0
			}else if mp[k-a%k] < mp[a] && mp[k-a%k] > 0{
				mp[a] -= mp[k-a%k]
				mp[k-a%k] = 0
			}
		}

	}
	for a,b := range mp{
		if a != 0 && b != 0 {
			return false
		}
	}

	return true

}

func main() {
	fmt.Println(canArrange([]int{-4,-7,5,2,9,1,10,4,-8,-3},3))

	
}
