package main


func ReverseList(head *ListNode) *ListNode {


	var pre  *ListNode
	for head != nil && head.Next != nil{
		Next := head.Next
		head.Next = pre
		head = Next
		pre = head
	}
	head.Next = pre
	return head
}
func main() {
	
}
