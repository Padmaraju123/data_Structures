package main

import (
	"fmt"
)

type node struct {
	number int
	prev   *node
	next   *node
}

type doublylinkedlist struct {
	len  int
	head *node
	tail *node
}

// insert value at last
func (dll *doublylinkedlist) InsertAtlast(num int) {
	var temp = &node{}
	temp.number = num
	temp.prev = nil
	temp.next = nil
	if dll.head == nil && dll.tail == nil {
		dll.head = temp
		dll.tail = temp
	} else {
		temp.prev = dll.tail
		dll.tail.next = temp
		dll.tail = temp
	}
	dll.len++
}

// Display/Traverse
func (dll *doublylinkedlist) Display() {
	var p *node = dll.head
	for p != nil {
		fmt.Printf("%d ->", p.number)
		p = p.next
	}
	fmt.Println()
}

// Display Reverse order
func (dll *doublylinkedlist) DisplayReverse() {
	var p *node = dll.tail
	for p != nil {
		fmt.Printf("%d ->", p.number)
		p = p.prev
	}
	fmt.Println()
}

// Insert at begining
func (dll *doublylinkedlist) InsertAtBegining(num int) {
	var temp = &node{}
	temp.number = num
	temp.prev = nil
	temp.next = nil
	if dll.head == nil && dll.tail == nil {
		dll.head = temp
		dll.tail = temp
	} else {
		temp.next = dll.head
		dll.head.prev = temp
		dll.head = temp
	}
	dll.len++
}

// delete from Begining
func (dll *doublylinkedlist) DeleteFromBeginging() {
	if dll.head == dll.tail {
		dll.head = nil
		dll.tail = nil
	} else {
		dll.head = dll.head.next
		dll.head.prev = nil
	}
	dll.len--
}

// delete form end of linkedlist
func (dll *doublylinkedlist) DeleteFromEnd() {
	if dll.head == dll.tail {
		dll.head = nil
		dll.tail = nil
	} else {
		dll.tail = dll.tail.prev
		dll.tail.next = nil
	}
	dll.len--
}

// Delete Specific index
func (dll *doublylinkedlist) DeleteFromSpecific(position int) {
	if dll.len >= position {
		var p *node = dll.head
		for i := 0; i < position-1; i++ {
			p = p.next
		}
	}
}


func main(){
	obj := doublylinkedlist{}
	
	sli := []int{10,20,30,40,50}

	for _,v := range sli{
		obj.InsertAtlast(v)
	}

	obj.Display()

	obj.InsertAtBegining(200)

	obj.Display()

	obj.DisplayReverse()

	obj.DeleteFromBeginging()

	obj.Display()

	obj.DeleteFromSpecific(2)

	obj.Display()

	obj.DeleteFromEnd()

	obj.Display()
}
