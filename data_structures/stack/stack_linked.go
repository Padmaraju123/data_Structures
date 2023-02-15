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
	size int
}

func (obj *List) length() int {
	return obj.size
}

func (obj *List) pushing(val int){
	newnode := Node{element: val}
	if obj.length() == 0 {
		obj.head = &newnode
	}else{
		newnode.next = obj.head
		obj.head = &newnode
	}
	obj.size++
}

func (obj *List)poping(){
	if obj.length()==0{
		fmt.Println("no nodes are found")
	}else{
		comm := obj.head
		for comm!=nil{
			fmt.Println(comm.element)
			comm = comm.next
		}
	}
}

func main() {
	obj := List{}
	reader := bufio.NewReader(os.Stdin)
	sli_b, _, _ := reader.ReadLine()
	sent := string(sli_b)

	split_S := strings.Split(sent, " ")

	for _, v := range split_S {
		vv, _ := strconv.Atoi(v)
		obj.pushing(vv)
	}

	fmt.Print("printing the values from latest to old\n")
	obj.poping()



}
