package main

import (
	"fmt"
)
type ListNode struct {
	Val int
	Next *ListNode
}
func deleteDuplicates(head *ListNode) *ListNode {

	pre := &ListNode{
		Val: -1,
		Next: head,
	}
	point := pre
	count := 0
	for head != nil && head.Next!=nil {
		for head.Next != nil && head.Val == head.Next.Val {
			count++
			head = head.Next
		}
		if count > 0  {
			point.Next = head.Next
			head = head.Next
			count = 0
		} else {
			point = head
			head = head.Next
		}
	}
	return pre.Next
}
func main() {
	d:= ListNode{
		Val:  2,
		Next: nil,
	}
	c:= ListNode{
		Val:  2,
		Next: &d,
	}
	b := ListNode{
		Val:  2,
		Next: &c,
	}
	a:= ListNode{
		Val:  1,
		Next: &b,
	}
	fmt.Println(deleteDuplicates(&a))
	
}
