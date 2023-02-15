//Input: linked list = 10->20->30->40->50->60, k = 4
//Output: 50->60->10->20->30->40.
//Explanation: k is smaller than the count of nodes in a linked list so (k+1 )th node i.e. 50 becomes the head node and 60’s next points to 10

//Input: linked list = 30->40->50->60, k = 2
//Output: 50->60->30->40.

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

func (obj *List) rotating(pos int) {

	comm := obj.head

	for i := 0; i < pos-1; i++ {
		comm = comm.next
	}

	kth := comm

	for comm.next != nil {
		comm = comm.next
	}

	comm.next = obj.head
	obj.head = kth.next
	kth.next = nil

}

func (obj *List) printing() {

	comm := obj.head
	for comm != nil {
		fmt.Print(comm.element, "-->")
		comm = comm.next
	}
}

func main() {
	obj := List{}

	sli := []int{10, 20, 30, 40, 50, 60}

	for _, v := range sli {
		obj.appending(v)
	}

	obj.rotating(4)

	obj.printing()

}
