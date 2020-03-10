package main

func reversePrint(head *ListNode) []int {
	temp  := []int{}
	for head!=nil{
		temp = append(temp, head.Val)
		head = head.Next
	}
	res := make([]int,len(temp))
	for i:= len(temp); i>=0 ; i--{

		res[len(temp) - i ] = temp[i]
	}
	return res



}
func main() {
	
}
