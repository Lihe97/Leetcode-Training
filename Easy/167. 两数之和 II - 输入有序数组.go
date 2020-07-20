package main

func twoSum(numbers []int, target int) []int {

	mp := map[int]int{}
	res := []int{}

	for i := 0 ; i < len(numbers) ; i ++{
		if _,ok := mp[target-numbers[i]];ok{
			res = append(res,mp[target-numbers[i]]+1)
			res = append(res,i+1)
			break
		}else{
			mp[numbers[i]] = i
		}
	}
	return res

}
func main() {
	
}
