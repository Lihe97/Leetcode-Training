package main
func removeDuplicateNodes(head *ListNode) *ListNode {

	mp := map[int]bool{}

	prev := &ListNode{}
	prev.Next = head

	mp[head.Val] = true
	head = head.Next
	pre := prev.Next

	for head != nil{
		if ! mp[head.Val]{
			mp[head.Val] = true
			head = head.Next
			pre = pre.Next
		}else{
			next := head.Next
			pre.Next = next
			head = next
		}
	}
	return prev.Next

}

func main() {
	
}
