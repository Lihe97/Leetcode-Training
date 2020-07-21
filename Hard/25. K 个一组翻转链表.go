package main
type ListNode struct {
	Val int
	Next *ListNode
}


//func reverseKGroup(head *ListNode, k int) *ListNode {
//
//	prev := &ListNode{-1,head}
//	pre , end := prev,prev
//	for end.Next != nil{
//
//		for i := 0 ; i < k && end != nil; i++{
//			end = end.Next
//		}
//		if end == nil{
//			break
//		}
//		start := pre.Next
//		temp := pre.Next
//		NNext := end.Next
//		pree := new(ListNode)
//		pree = nil
//		for i := 0 ; i < k ; i++{
//			Next := start.Next
//			start.Next = pree
//			pree = start
//			start = Next
//		}
//		pre.Next = end
//		temp.Next = NNext
//		pre = temp
//		end = pre
//
//	}
//	return prev.Next
//
//}
func reverseKGroup(head *ListNode, k int) *ListNode {

	prev := &ListNode{}
	prev.Next = head

	cur := prev
	curr := prev

	for  curr != nil {

		cur = curr
		for i := 0 ; i < k  && cur != nil ; i ++{
			cur = cur.Next
		}
		if cur == nil{
			break
		}
		a := curr.Next
		b := a.Next
		for i := 0 ; i < k-1 ; i ++{
			c := b.Next
			b.Next = a
			a = b
			b = c
		}
		c := curr.Next
		c.Next = b
		curr.Next = a

		curr = c

	}
	return prev.Next

}
func main() {

	
}
