package main

import (
	"fmt"
)

type Hashchain struct{
	hashtable_size int
	hashtable []int
}

type Node struct {
	element int
	next    *Node
}

type List struct {
	head *Node
	tail *Node
	size int
}

func (obj *List) isempty() bool {
	if obj.size == 0 {
		return true
	}
	return false
}

func (obj *List) length() int {
	return obj.size
}

func (obj *List) display() {
	comm := obj.head
	for comm != nil {
		fmt.Print(comm.element, "-->")
		comm = comm.next
	}
	fmt.Println()
}

func (obj *List) addlast(val int) {
	newnode := Node{element: val}
	if obj.isempty() {
		obj.head = &newnode
	} else {
		obj.tail.next = &newnode
	}
	obj.tail = &newnode
	obj.size++
}

func (obj *List) addfirst(val int) {
	newnode := Node{element: val}
	if obj.isempty() {
		obj.head = &newnode
		obj.tail = &newnode
	} else {
		newnode.next, obj.head = obj.head, &newnode
	}
	obj.size++
}

func (obj *List) addany(val, pos int) {
	newnode := Node{element: val}
	comm := obj.head
	for i := 0; i < pos-1; i++ {
		comm = comm.next
	}
	temp := comm.next
	comm.next = &newnode
	newnode.next = temp
	obj.size++
}

func (obj *List) removefirst() int {
	if obj.isempty() {
		fmt.Println("no nodes are found")
		return 0
	}

	comm := obj.head
	val := obj.head.element

	obj.head = comm.next
	obj.size--

	return val

}

func (obj *List) removelast() int {
	if obj.isempty() {
		fmt.Println("no nodes are found to delete")
		return 0
	}

	comm := obj.head
	for i := 1; i < obj.length()-1; i++ {
		comm = comm.next
	}
	value := comm.next.element
	temp := comm
	comm.next = nil
	obj.tail = temp
	obj.size--
	return value

}

func (obj *List) removeany(pos int) int {
	if obj.isempty() {
		fmt.Println("no nodes are found")
		return 0
	}

	comm := obj.head
	for i := 1; i < pos-1; i++ {
		comm = comm.next
	}

	val := comm.next.element
	comm.next = comm.next.next
	obj.size--
	return val
}

func (obj *List) search(val int) int {
	if obj.isempty() {
		fmt.Println("no nodes are found to search")
		return -1
	}

	comm := obj.head
	i := 0
	for comm != nil {
		if comm.element == val {
			return i
		}
		comm = comm.next
		i++
	}

	return -1
}

func (obj *List) insertsorted(val int) {
	newnode := Node{element: val}
	if obj.isempty() {
		obj.head = &newnode
	} else {
		p := obj.head
		q := obj.head
		for p != nil && p.element < val {
			q = p
			p = p.next
		}
		if p == obj.head {
			newnode.next = obj.head
			obj.head = &newnode
		} else {
			newnode.next = q.next
			q.next = &newnode
		}
	}
	obj.size++
}

func (obj *)

func main() {
	obj := Hashchain{}

	obj.hashtable_size = 10
	obj.hashtable = make([]int,obj.hashtable_size)

	for i:=0;i<obj.hashtable_size;i++{
		obj.hashtable[i] = List{}

	}

}
