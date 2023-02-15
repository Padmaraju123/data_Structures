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

type List struct{
	head *Node
	tail *Node
	len int
}

func (obj *List)length()int{
	return obj.len
}

func (obj *List)isempty()bool{
	if obj.length()==0{
		return true
	}
	return false
}

func (obj *List)Enquening(val int){
	newnode := Node{element: val}
	if obj.isempty(){
		obj.head = &newnode
	}else{
		obj.tail.next = &newnode
	}
	obj.tail = &newnode
	obj.len++
}

func (obj *List)Dequening(){
	if obj.isempty(){
		fmt.Println("no nodes are found to deque in List.....")
		return
	}
	comm := obj.head
	for comm!=nil{
		fmt.Print(comm.element,"--->")
		comm = comm.next
	}
}



func main() {
	fmt.Println("enter the number line with spaces...")
	reader := bufio.NewReader(os.Stdin)
	sli_c, _, _ := reader.ReadLine()
	sent := string(sli_c)

	split_c := strings.Split(sent, " ")

	obj := List{}

	for _, v := range split_c {
		kk, _ := strconv.Atoi(v)
		obj.Enquening(kk)
	}

	obj.Dequening()


}
