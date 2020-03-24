package main

import "fmt"
type ListNode struct {
	Val int
	Next *ListNode
}
func insertionSortList(head *ListNode) *ListNode {

	res := &ListNode{
		Val:  11,
		Next: nil,
	}
	cur := res

	cur.Next = head
	head = head.Next
	cur.Next.Next = nil

	for head!= nil{
		cur = res
		//fmt.Println("cur",cur)
		//fmt.Println()
		for cur.Next != nil && head != nil{
			//fmt.Println("head",head)
			//fmt.Println("cur.Next",cur.Next)
			if head.Val <= cur.Next.Val{
				Next := head.Next
				head.Next = cur.Next
				cur.Next = head
				head = Next
				break
			}else{
				cur = cur.Next
			}
		}
		if cur.Next == nil{
			Next := head.Next
			cur.Next = head
			head.Next = nil
			head = Next
			//continue
		}

	}
	//for res.Next != nil{
	//	fmt.Println(res.Next)
	//	res.Next = res.Next.Next
	//}
	return res.Next
}
func main() {

	e := &ListNode{
		Val:  0,
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
		Val:  5,
		Next: c,
	}

	a := &ListNode{
		Val:  1,
		Next: b,
	}
	fmt.Println(insertionSortList(a))
}
