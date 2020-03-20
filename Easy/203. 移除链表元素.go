package main
type ListNode struct {
	Val int
	Next *ListNode
}
func removeElements(head *ListNode, val int) *ListNode {

	prev := &ListNode{0,head}
	if head == nil{
		return head
	}
	cur := prev
	temp := head

	for temp != nil{
		if temp.Val == val{
			cur.Next = temp.Next
			temp = temp.Next
		}else{
			cur.Next = temp
			cur = cur.Next
			temp = temp.Next
		}


	}
	return prev.Next
}
func main() {
	
}
