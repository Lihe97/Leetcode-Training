package main
type ListNode struct {
	Val int
	Next *ListNode
}
func deleteNode(head *ListNode, val int) *ListNode {

	prev := &ListNode{
		Val:  0,
		Next: head,
	}
	cur := prev
	for head != nil && head.Val != val{
		head = head.Next
		cur = cur.Next
	}
	cur.Next = head.Next
	return prev.Next

}
func main() {
	
}
