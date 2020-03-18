package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	if l1 == nil{
		return l2
	}
	if l2 == nil{
		return l1
	}
	if l2 == nil && l1 ==nil{
		return res
	}
	if l1.Val<= l2.Val{
		res = l1
		l1 = l1.Next
	}else {
		res = l2
		l2 = l2.Next
	}
	back := res
	for l1 != nil && l2 != nil{
		if l1.Val <= l2.Val{
			back.Next = l1
			back = l1
			l1 = l1.Next
		}else {
			back.Next = l2
			back = l2
			l2 = l2.Next
		}
	}
	if l1 != nil{
		back.Next = l1
		back = l1
		l1 = l1.Next
	}
	if l2 != nil{
		back.Next = l2
		back = l2
		l2 = l2.Next
	}
	return res

}
func main() {
	
}
