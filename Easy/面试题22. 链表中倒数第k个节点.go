package main

func getKthFromEnd(head *ListNode, k int) *ListNode {
	prev := &ListNode{
		Val:  0,
		Next: head,
	}
	first := prev
	second := prev
	for i := 0 ; i < k ; i++{
		first = first.Next
	}
	for first!=nil{
		second = second.Next
		first = first.Next
	}
	return second
}
func main() {
	
}
