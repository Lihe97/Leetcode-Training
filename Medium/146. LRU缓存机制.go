package main

import (
	"fmt"
)
type DubleNode struct{
	key int
	val int
	prev *DubleNode
	next *DubleNode
}

type LRUCache struct {
	capacity int
	cache map[int]*DubleNode
	head,end *DubleNode
}


func Constructor(capacity int) LRUCache {

	end := &DubleNode{}
	head := &DubleNode{}
	head.next = end
	end.prev = head

	return LRUCache{
		capacity: capacity,
		cache: map[int]*DubleNode{},
		head:     head,
		end:     end,
	}
}


func (this *LRUCache) Get(key int) int {


	if this.cache[key] == nil{
		return -1
	}
	node := this.cache[key]


	this.removeNode(node)
	this.insertHead(node)

	return node.val
}
func(this *LRUCache) removeNode(node *DubleNode){
	prev,next := node.prev,node.next
	prev.next = next
	next.prev = prev
}
func(this *LRUCache) insertHead(node *DubleNode){
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}


func (this *LRUCache) Put(key int, value int)  {

	if this.cache[key] == nil{
		node := &DubleNode{
			key : key,
			val:  value,
			prev: nil,
			next: nil,
		}
		if this.capacity == len(this.cache){

			delete(this.cache,this.end.prev.key)

			this.removeNode(this.end.prev)
			this.insertHead(node)
		}else{
			this.insertHead(node)
		}
		this.cache[key] = node
	}else{
		node := this.cache[key]
		node.val = value
		this.removeNode(node)
		this.insertHead(node)
	}


}

func main() {

	obj := Constructor(1)

	obj.Put(2,1)

	obj.Put(3,2)

	fmt.Println("拿2",obj.Get(2))


}
