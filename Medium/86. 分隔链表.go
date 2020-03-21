package main

import (
	"fmt"
)

func partition(head *ListNode, x int) *ListNode {

	prev := &ListNode{
		Val:  1,
		Next: nil,
	}
	prev2 := &ListNode{
		Val:  1,
		Next: nil,
	}
	cur1 := prev
	cur2 := prev2
	for head!=nil{
		if head.Val < x{
			cur1.Next = head
			cur1 = head

		}else{
			cur2.Next = head
			cur2 = head

		}
		head = head.Next
	}
	cur2.Next = nil
	cur1.Next = prev2.Next
	//for prev.Next!=nil{
	//	fmt.Println(prev.Next)
	//	prev.Next = prev.Next.Next
	//	time.Sleep(2*time.Second)
	//}
	return prev.Next

}
func main() {
	f:= ListNode{
		Val:  2,
		Next: nil,
	}
	e := ListNode{
		Val:  5,
		Next: &f,
	}
	d:= ListNode{
		Val:  2,
		Next: &e,
	}
	c:= ListNode{
		Val:  3,
		Next: &d,
	}
	b := ListNode{
		Val:  4,
		Next: &c,
	}
	a:= ListNode{
		Val:  1,
		Next: &b,
	}
	fmt.Println(partition(&a,3))
	
}
