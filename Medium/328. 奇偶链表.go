package main

import (
	"fmt"
)

func oddEvenList(head *ListNode) *ListNode {

	if head == nil{
		return nil
	}
	if head != nil && head.Next == nil{
		return head
	}
	if head != nil && head.Next != nil && head.Next.Next == nil{
		return head
	}
	prev1 := &ListNode{
		Val:  0,
		Next: head,
	}
	prev2 := &ListNode{
		Val:  0,
		Next: head.Next,
	}
	cur1 := head
	cur2 := head.Next
	count := 1
	temp := head.Next.Next
	//fmt.Println("1111temp",temp)


	for temp != nil{
		if count % 2 == 1{
			cur1.Next = temp
			cur1 = temp
			//cur1.Next = nil
		}else{
			cur2.Next = temp
			cur2 = temp
			//cur2.Next = nil
		}
		count ++
		//fmt.Println("temp",temp)
		temp = temp.Next
	}
	//fmt.Println("----1",cur1)
	//fmt.Println("----2",cur2)
	cur2.Next = nil
	cur1.Next = prev2.Next
	//for prev1.Next != nil{
	//	fmt.Println(prev1.Next)
	//	time.Sleep(2*time.Second)
	//	prev1.Next = prev1.Next.Next
	//}
	return prev1.Next
}
func main() {

	e := &ListNode{
		Val:  5,
		Next: nil,
	}
	d := &ListNode{
		Val:  4,
		Next: e,
	}
	c := &ListNode{
		Val:  3,
		Next: d,
	}
	b := &ListNode{
		Val:  2,
		Next: c,
	}
	a := &ListNode{
		Val:  1,
		Next: b,
	}
	fmt.Println(oddEvenList(a))
}
