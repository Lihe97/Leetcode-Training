package main

type ListNode struct {
	Val int
	Next *ListNode
}
func mergeKLists(lists []*ListNode) *ListNode {

	if len(lists) == 0{
		return nil
	}
	res := lists[0]

	for i := 1 ; i < len(lists)  ; i++{
		if lists[i] == nil{
			continue
		}
		if res == nil{
			res = lists[i]
			continue
		}

		res = mergeTwo(res,lists[i])

	}
	return res

}
func mergeTwo(list1 *ListNode,list2 *ListNode) *ListNode {

	temp := &ListNode{}
	if list1.Val > list2.Val{
		temp = list2
		list2 = list2.Next
	}else {
		temp = list1
		list1 = list1.Next
	}
	res := temp
	for list1 != nil && list2 != nil{
		if list1.Val > list2.Val{
			temp.Next = list2
			temp = temp.Next
			list2 = list2.Next

		}else {
			temp.Next = list1
			temp = temp.Next
			list1 = list1.Next
		}
	}
	if list1 != nil{
		temp.Next = list1
	}
	if list2 != nil{
		temp.Next = list2
	}
	return res
}
func main() {
	
}
