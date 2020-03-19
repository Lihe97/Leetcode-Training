package main

import "fmt"

func kthToLast(head *ListNode, k int) int {

	first := head
	second := head
	for i:= 0 ; i < k ; i ++{
		first = first.Next
	}
	for first!=nil{
		first = first.Next
		second = second.Next

	}
	return second.Val

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


	fmt.Println(kthToLast(&a,3))
}
