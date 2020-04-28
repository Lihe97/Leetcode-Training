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

	end := &DubleNode{
		key : 0,
		val:  0,
		prev: nil,
		next: nil,
	}
	head := &DubleNode{
		key: 0,
		val:  0,
		prev: nil,
		next: end,
	}
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

	prev,next := node.prev,node.next
	prev.next = next
	next.prev = prev

	this.insertHead(node)

	return node.val
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
			//this.cache[this.end.prev.val] = nil
			delete(this.cache,this.end.prev.key)
			fmt.Println(this.cache)
			this.end.prev.prev.next = this.end
			this.end.prev = this.end.prev.prev

			this.insertHead(node)
		}else{
			this.insertHead(node)
		}
		this.cache[key] = node
	}else{
		node := this.cache[key]
		node.val = value

		node.prev.next = node.next
		node.next.prev = node.prev
		this.insertHead(node)
	}


}

func main() {

	obj := Constructor(1)

	obj.Put(2,1)

	obj.Put(3,2)

	fmt.Println("拿2",obj.Get(2))


}
