package main
func deleteDuplicates(head *ListNode) *ListNode {

	temp := head
	res  := head
	count := 1
	for temp.Next !=nil && temp!=nil{
		if temp.Next.Val == temp.Val{
			temp = temp.Next
			count ++
		}else{
			if count == 1 {
				res.Next = temp
				res = res.Next
				temp = temp.Next
			}else{
				count = 1
				temp = temp.Next
			}
		}

	}
	return head.Next
}

func main() {
	
}
