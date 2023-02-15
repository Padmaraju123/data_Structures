/*Input: linked list = 10->20->30->40->50->60, k = 4
Output: 50->60->10->20->30->40.
Explanation: k is smaller than the count of nodes in a linked list so (k+1 )th node i.e. 50 becomes the head node and 60’s next points to 10

Input: linked list = 30->40->50->60, k = 2
Output: 50->60->30->40. */

package main

import "fmt"

type Node struct {
	element int
	next    *Node
}

type List struct {
	head *Node
	tail *Node
	len  int
}

func (obj *List) length() int {
	return obj.len
}

func (obj *List) isempty() bool {
	if obj.len == 0 {
		return true
	}
	return false
}

func (obj *List) appending(val int) {
	newnode := Node{element: val}
	if obj.isempty() {
		obj.head = &newnode
	} else {
		obj.tail.next = &newnode
	}
	obj.tail = &newnode
	obj.len++
}

func (obj *List) printing() {
	if obj.isempty() {
		fmt.Println("no nodes are there to print..")
	} else {
		comm := obj.head
		for comm != nil {
			fmt.Print(comm.element, "-->")
			comm = comm.next
		}
	}
	fmt.Println()
}

func (obj *List)rotating(k int){
	comm := obj.head

	for i:=0 ; i<k-1;i++{
		comm = comm.next
	}

	add1 := comm

	for comm.next != nil{
		comm =comm.next
	}

	comm.next = obj.head
	obj.head = add1.next
	add1.next = nil



}

func main() {

	obj := List{}

	fmt.Println("enter the size of slice...")
	var size int
	fmt.Scanln(&size)

	sli := make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Scan(&sli[i])
		obj.appending(sli[i])

	}

	obj.printing()

	fmt.Println("enter the postion value to rotate...")

	var pos int
	fmt.Scan(&pos)


	obj.rotating(pos)

	obj.printing()



}
