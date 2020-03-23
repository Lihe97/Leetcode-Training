	package main

type listNode struct {
      Val int
      Next *listNode
 }
func middleNode(head *listNode) *listNode {
	num := 1
	hashmap := map[int]*listNode{}
	for head!=nil{
		hashmap[num] = head
		head = head.Next
		num++
	}
	//fmt.Println(num)
	if num%2 == 0{
		return hashmap[num/2]
	}else{
		return  hashmap[num/2+1]
	}
	//return
}
func middleNode1(head *listNode) *listNode {
	slow , fast := head ,head
	for fast != nil && fast.Next!=nil{
		slow ,fast = slow.Next,fast.Next.Next
	}
	return slow
}
func main() {
	
}
