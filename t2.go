package main

type ListNode struct {
	Val int
	Next *ListNode
}
func nodesBetweenCriticalPoints(head *ListNode) []int {
	nums := []int{}
	for head != nil{
		nums = append(nums,head.Val)
		head = head.Next
	}
	pos := []int{}

	for i := 1 ; i < len(nums) -1; i ++{
		if nums[i] > nums[i-1] && nums[i] > nums[i+1]{
			pos = append(pos,i)
		}
		if nums[i] < nums[i-1] && nums[i] < nums[i+1]{
			pos = append(pos,i)
		}
	}
	if len(pos) == 0 || len(pos) == 1{
		return []int{-1,-1}
	}
	max := pos[len(pos)-1] - pos[0]
	min := 1<<32
	for i := 1 ; i < len(pos) ; i ++{
		if pos[i] - pos[i-1] < min{
			min = pos[i] - pos[i-1]
		}
	}
	res := []int{min,max}
	return res
}

func main()  {
}
