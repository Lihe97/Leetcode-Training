package main

import "fmt"


func rotateRight(head *ListNode, k int) *ListNode {
	prev := &ListNode{
		Val:  0,
		Next: head,
	}
	length := 0
	for head!=nil{
		length ++
		head = head.Next
	}
	if length == 0{
		return nil
	}
	head = prev.Next
	if k ==0{
		return head
	}

	k = k%length

	first := prev
	second := prev
	for i:= 0 ; i < k ; i++{
		first = first.Next
	}
	for first.Next!=nil{
		first = first.Next
		second = second.Next
	}
	res := second.Next
	first.Next = prev.Next
	second.Next = nil
	//for res!= nil{
	//	fmt.Println(res)
	//	res = res.Next
	//}
	return  res

}
func main() {
	//e := ListNode{
	//	Val:  5,
	//	Next: nil,
	//}
	//d:= ListNode{
	//	Val:  4,
	//	Next: &e,
	//}
	//c:= ListNode{
	//	Val:  3,
	//	Next: &d,
	//}
	//b := ListNode{
	//	Val:  2,
	//	Next: &c,
	//}
	a:= ListNode{
		Val:  1,
		Next: nil,
	}
	fmt.Println(rotateRight(&a,1))
}
