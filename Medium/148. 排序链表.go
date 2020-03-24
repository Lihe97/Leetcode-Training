package main
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func sortList(head *ListNode) *ListNode{
	if head == nil || head.Next == nil{
		return head
	}
	slow , fast := head,head
	for fast.Next != nil && fast.Next.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
	}
	right := sortList(slow.Next)
	slow.Next = nil
	left := sortList(head)
	return mergeTwoList(left,right)
}
func mergeTwoList(a,b *ListNode) *ListNode{

	prev := new(ListNode)
	cur := prev
	for a != nil && b != nil{
		if a.Val <= b.Val{
			cur.Next = a
			cur = cur.Next
			a = a.Next
		}else{
			cur.Next = b
			cur = cur.Next
			b = b.Next
		}
	}
	if a != nil{
		cur.Next = a
	}
	if b != nil{
		cur.Next = b
	}
	return prev.Next
}
func main() {
	
}
