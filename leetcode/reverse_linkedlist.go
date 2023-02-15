/*Given a pointer to the head node of a linked list, the task is to reverse the linked list. We need to reverse the list by changing the links between nodes.

Examples:

Input: Head of following linked list
1->2->3->4->NULL
Output: Linked list should be changed to,
4->3->2->1->NULL

Input: Head of following linked list
1->2->3->4->5->NULL
Output: Linked list should be changed to,
5->4->3->2->1->NULL */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	val  int
	next *Node
}

type List struct {
	head *Node
	tail *Node
}

// this will gives the actual length of the linked list..
func (obj *List) length() int {
	c := obj.head
	var le int
	for c != nil {
		le++
		c = c.next
	}

	return le
}

// this gives the bool value whether the list is empty or not
func (obj *List) isempty() bool {
	if obj.length() == 0 {
		return true
	}
	return false
}

// appending the nodes to the list....
func (obj *List) appending(v int) {
	newnode := Node{val: v}
	if obj.isempty() {
		obj.head = &newnode
	} else {
		obj.tail.next = &newnode
	}
	obj.tail = &newnode
}

// printing the values in the linked list....
func (obj *List) printing() {
	comm := obj.head
	for comm != nil {
		fmt.Print(comm.val, "--->")
		comm = comm.next
	}
	fmt.Println()
}

// Reverse the linked list
func reverseList(head *Node) *Node {
	var prev,curr, nn *Node
	curr = head
	prev = nil
	for curr != nil{
		nn = curr.next
		curr.next = prev
		prev = curr
		curr = nn

	}
	return prev
}

// printing the values of reverse in the linked list....
func printing(troot *Node) {
	comm := troot
	for comm != nil {
		fmt.Print(comm.val, "--->")
		comm = comm.next
	}
	fmt.Println()
}


// this is main function where the exection starts...
func main() {
	fmt.Println("enter the number line with spaces....")

	obj := List{}

	reader := bufio.NewReader(os.Stdin)
	sli_c, _, _ := reader.ReadLine()

	snt := string(sli_c)

	split_c := strings.Split(snt, " ")

	for _, v := range split_c {
		kk, _ := strconv.Atoi(v)
		obj.appending(kk)

	}

	obj.printing()

	out := reverseList(obj.head)

	printing(out)

}
