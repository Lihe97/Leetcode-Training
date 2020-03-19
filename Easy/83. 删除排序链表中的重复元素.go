package main

import "fmt"
type ListNode struct {
	Val int
	Next *ListNode
}
func deleteDuplicates(head *ListNode) *ListNode {

	prev := &ListNode{0, head}
	cur := prev
	count := 0
	for(head != nil && head.Next != nil) {
		for (head.Next != nil && head.Val == head.Next.Val ){
			count++
			head=head.Next
		}
		if count > 0 {
			cur.Next = head.Next
			head=head.Next
			count = 0
		}else{
			cur.Next = head
			cur = cur.Next
			head = head.Next
		}
	}
	return prev.Next
}
func main() {
	d:= ListNode{
		Val:  4,
		Next: nil,
	}
	c:= ListNode{
		Val:  3,
		Next: &d,
	}
	b := ListNode{
		Val:  2,
		Next: &c,
	}
	a:= ListNode{
		Val:  2,
		Next: &b,
	}
	fmt.Println(deleteDuplicates(&a))
}
