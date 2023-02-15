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

func (obj *List) middleval() int {
	comm := obj.head

	if obj.len%2 != 0 {
		for i := 0; i < obj.len/2; i++ {
			comm = comm.next
		}
		return comm.element
	}

	for j := 0; j < obj.len/2; j++ {
		comm = comm.next
	}

	return comm.element

}

func main() {
	obj := List{}

	sli := []int{1, 2, 3, 4, 5,6,7}

	for _, v := range sli {
		obj.appending(v)
	}

	fmt.Println(obj.middleval())
}
