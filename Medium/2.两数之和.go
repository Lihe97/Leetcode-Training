package main

import (
	"fmt"
)

//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//
//如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
//您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//示例：
//
//输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807
//
//模拟加法，自己写的代码太乱，把解答的放上来了
//有很多坑，如[1],[9,9,9,9,9]这种进位问题
type ListNode struct {
	Val int
	Next *ListNode
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	cur := res
	num := 0
	for l1 != nil || l2 != nil || num > 0{
		sum:=0
		if l1!=nil{
			sum += l1.Val
			l1 = l1.Next
		}
		if l2!=nil{
			sum  += l2.Val
			l2 = l2.Next
		}
		sum += num
		num = sum/10
		temp := &ListNode{
			Val:  sum%10,
			Next: nil,
		}
		cur.Next = temp
		cur = cur.Next

	}
	return res.Next

}
func main() {
	a := new(ListNode)
	b := new(ListNode)
	c := new(ListNode)
	a.Val = 1
	b.Val = 2
	c.Val = 3
	b.Next = c
	fmt.Println(addTwoNumbers(a,b))
	
}
