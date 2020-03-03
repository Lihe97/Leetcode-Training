package main

type ListNode struct {
     Val int
      Next *ListNode
}
func reverseList(head *ListNode) *ListNode {

	if head == nil {
		return head
	}
	var pre,next *ListNode
	for head!=nil{
		next = head.Next
		head.Next = pre
		pre = head
		head = next

	}
	return pre
}
func main() {
	
}
