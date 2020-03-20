package main

import (
	"fmt"
)
type ListNode struct {
	Val int
	Next *ListNode
}
func swapPairs(head *ListNode) *ListNode {
	prev := &ListNode{
		Val:  0,
		Next: head,
	}
	cur := prev
	temp := head
	for head != nil && head.Next!=nil{
		head = head.Next.Next
		cur.Next = temp.Next
		cur.Next.Next = temp
		cur = temp
		cur.Next = nil
		temp  = head
	}
	if head != nil && head.Next == nil{
		cur.Next = head
	}
	//fmt.Println(prev.Next)
	//fmt.Println(prev.Next.Next)
	//fmt.Println(prev.Next.Next.Next)

	//for prev.Next != nil{
	//	fmt.Println(prev.Next)
	//	prev.Next = prev.Next.Next
	//	time.Sleep(2*time.Second)
	//}
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
		Val:  1,
		Next: &b,
	}
	fmt.Println(swapPairs(&a))
	
}
