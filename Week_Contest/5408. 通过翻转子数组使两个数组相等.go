package main

func canBeEqual(target []int, arr []int) bool {

	mp := map[int]int{}

	for i := 0 ; i < len(arr) ; i ++{
		mp[target[i]]++
		mp[arr[i]] --
	}
	for _,v := range mp{
		if v != 0{
			return false
		}
	}
	return true
}
func main() {
	
}
