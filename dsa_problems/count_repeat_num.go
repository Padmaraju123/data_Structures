package main

import "fmt"

type Node struct {
	element int
	next    *Node
}

type List struct {
	head *Node
	len  int
	tail *Node
}

func (obj *List) appending(val int) {
	newnode := Node{element: val}
	if obj.len == 0 {
		obj.head = &newnode
	} else {
		obj.tail.next = &newnode
	}
	obj.tail = &newnode
	obj.len++
}

var c int

func (obj *List)repeating_val(comm *Node,val int)int{

	if comm==nil{
		return c
	}
	
	if val == comm.element{
		c++
	}
	
	return obj.repeating_val(comm.next,val)

}

func main() {
	obj := List{}

	sli := []int{1, 2, 3, 4, 5,6,7,3,3,2,2}

	for _, v := range sli {
		obj.appending(v)
	}

	fmt.Println(obj.repeating_val(obj.head,3))
}