package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	prev    *Node
	element int
	next    *Node
}

type List struct {
	head *Node
	tail *Node
	len  int
}

// creating a double linked list with nodes
func (obj *List) appending(val int) {
	newnode := Node{element: val}
	if obj.len == 0 {
		obj.head = &newnode
	} else {
		obj.tail.next = &newnode
		newnode.prev = obj.tail
	}
	obj.tail = &newnode
	obj.len++
}

// printing the values in forward order of double linked list
func (obj *List) printing_for() {
	comm := obj.head
	for comm != nil {
		fmt.Print(comm.element, "-->")
		comm = comm.next
	}
	fmt.Println()
}

// printing the values in backward order of double linked list
func (obj *List) printing_bak() {
	comm := obj.tail
	for comm != nil {
		fmt.Print(comm.element, "-->")
		comm = comm.prev
	}
	fmt.Println()
}

func (obj *List) insert_beg(val int) {
	newnode := Node{element: val}
	if obj.len == 0 {
		obj.head = &newnode
		obj.tail = &newnode
	} else {
		temp := obj.head
		obj.head.prev = &newnode
		obj.head = &newnode
		newnode.next = temp
	}
	obj.len++
}

func (obj *List) deletion_beg() {
	if obj.len == 0 {
		fmt.Println("no nodes are found in double linked list....")
	} else {
		temp := obj.head.next
		obj.head.next.prev = nil
		obj.head = temp
		obj.len--
	}
}

func (obj *List) deletion_end() {
	if obj.len == 0 {
		fmt.Println("no nodes are found to delete....")
	} else {
		temp := obj.tail.prev
		obj.tail.prev.next = nil
		obj.tail = temp
		obj.len--
	}
}

func (obj *List) insertion_bw(pos, val int) {
	newnode := Node{element: val}
	if obj.len == 0 {
		obj.head = &newnode
		obj.tail = &newnode
	} else {
		comm := obj.head
		for i := 1; i < pos; i++ {
			comm = comm.next
		}

		add1 := comm

		add2 := comm.next

		comm.next = &newnode

		newnode.prev = add1

		newnode.next = add2

		add2.prev = &newnode

	}
	obj.len++
}

func (obj *List) deletion_bw(pos int) {
	if obj.len == 0 {
		fmt.Println("no nodes are there to delete...")
	} else {
		comm := obj.head
		for i := 1; i < pos; i++ {
			comm = comm.next
		}
		temp := comm
		comm.next = comm.next.next
		comm.next.next.prev = temp
		obj.len--
	}
}

func (obj *List) searching(val int) bool {
	if obj.len == 0 {
		fmt.Println("no nodes are there to search the given value...")
	} else {
		comm := obj.head
		for comm != nil {
			if comm.element == val {
				return true
			}
			comm = comm.next
		}

	}
	return false
}

func (obj *List)reverse_list(){
	if obj.len==0{
		fmt.Println("no nodes have to reverse it...")
	}else{
		obj.tail.next = obj.head
	}
}

func main() {
	obj := List{}

	reader := bufio.NewReader(os.Stdin)
	sli_c, _, _ := reader.ReadLine()
	sent := string(sli_c)

	split_c := strings.Split(sent, " ")

	for _, v := range split_c {
		k, _ := strconv.Atoi(v)
		obj.appending(k)
	}

	obj.printing_for()

	obj.printing_bak()

	obj.insert_beg(100)

	obj.printing_for()

	obj.deletion_beg()

	obj.printing_for()

	obj.deletion_end()

	obj.printing_for()

	obj.insertion_bw(2, 100)

	obj.printing_for()

	obj.deletion_bw(2)

	obj.printing_for()

	fmt.Println(obj.searching(3))

	obj.reverse_list()

	obj.printing_for()

}
