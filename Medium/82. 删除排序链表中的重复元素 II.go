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
		Val: 0,
		Next: head,
	}
	point := pre
	count := 1
	for head.Next != nil && head!=nil {
		for head.Next != nil && head.Val == head.Next.Val {
			count++
			head = head.Next
		}
		if count == 1 {
			point.Next = head.Next
			point = point.Next
			head = head.Next
		} else {
			head = head.Next
			count = 1
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
		Val:  1,
		Next: &c,
	}
	a:= ListNode{
		Val:  1,
		Next: &b,
	}
	fmt.Println(deleteDuplicates(&a))
	
}
