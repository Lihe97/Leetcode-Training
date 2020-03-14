package main

import "fmt"
type ListNode struct {
	Val int
	Next *ListNode
}
func kthToLast(head *ListNode, k int) int {
	cur := 1
	var res *ListNode
	temp := head
	for head!=nil{
		if k == cur{
			res = temp
		}
		if cur >k{
			res = res.Next
		}
		cur ++
		head= head.Next

	}
	return res.Val
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
