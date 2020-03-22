package main

type ListNode struct {
	Val int
	Next *ListNode
}
func insert(prev *ListNode,x int ,y int){
	temp := prev
	for i := 0 ; i < x ; i++{
		temp = temp.Next
	}
	node := &ListNode{y,temp.Next}
	temp.Next = node
}

func createTargetArray(nums []int, index []int) []int {
	prev := &ListNode{}

	for i:= 0 ; i < len(nums) ; i++{
		insert(prev,index[i],nums[i])
	}

	res := []int{}
	for prev.Next !=nil{
		res = append(res, prev.Next.Val)
		prev.Next = prev.Next.Next
	}
	return res
}
func main() {
	nums := []int{0,1,2,3,4}
	index := []int{0,1,2,2,1}
	createTargetArray(nums,index)
	
}
