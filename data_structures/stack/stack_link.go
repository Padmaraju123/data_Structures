package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	element int
	next    *Node
}

type List struct {
	head *Node
	tail *Node
	size int
}

func (obj *List) length() int {
	return obj.size
}

func (obj *List) pushing(val int) {
	newnode := Node{element: val}
	if obj.length() == 0 {
		obj.head = &newnode
	} else {
		newnode.next = obj.head
		obj.head = &newnode
	}
	obj.size++
}

func (obj *List) printing() {
	if obj.length() == 0 {
		fmt.Println("no elements are found....")
	} else {
		comm := obj.head
		for comm != nil {
			fmt.Print(comm.element, "-->")
			comm = comm.next
		}
		fmt.Println()
	}
}

func (obj *List) poping() {
	if obj.length() == 0 {
		fmt.Println("no elements are found to pop in linked list....")
	} else {
		comm := obj.head
		for comm != nil {
			fmt.Printf("%d\n", comm.element)
			comm = comm.next
		}
	}
}

func main() {

	obj := List{}

	fmt.Print("enter the number line with spaces\n")

	reader := bufio.NewReader(os.Stdin)
	sli_c, _, _ := reader.ReadLine()
	sent := string(sli_c)

	split_c := strings.Split(sent, " ")

	for _, v := range split_c {
		vv, _ := strconv.Atoi(v)
		obj.pushing(vv)
	}

	obj.printing()

	obj.poping()

}
