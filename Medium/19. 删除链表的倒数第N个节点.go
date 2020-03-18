package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	ans := &ListNode{0, head}
	for ; n > 0; n -- {
		head = head.Next
	}
	tmp := ans
	for head != nil{
		head = head.Next
		tmp = tmp.Next
	}
	tmp.Next = tmp.Next.Next
	return ans.Next
}
func main() {
	
}
