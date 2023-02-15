package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct{
	val int
	next *Node
}

type List struct{
	head *Node
	tail *Node
}


//this will gives the actual length of the linked list..
func (obj *List)length()int{
	c := obj.head
	var le int
	for c!=nil{
		le++
		c = c.next
	}

	return le
}


//this gives the bool value whether the list is empty or not
func (obj *List)isempty()bool{
	if obj.length()==0{
		return true
	}
	return false
}	


//appending the nodes to the list....
func (obj *List)appending(v int){
	newnode := Node{val: v}
	if obj.isempty(){
		obj.head = &newnode
	}else{
		obj.tail.next = &newnode
	}
	obj.tail = &newnode
}

//printing the values in the linked list....
func (obj *List)printing(){
	comm := obj.head
	for comm!=nil{
		fmt.Print(comm.val,"--->")
		comm = comm.next
	}
	fmt.Println()
}

//searching a given val return the index value of the linked list
func (obj *List)searching_val(num int)int{
	comm := obj.head
	var i int
	for comm!=nil{
		if comm.val==num{
			return i
		}
		i++
		comm = comm.next
	}
	return -1
}



//this is main function where the exection starts...
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

	fmt.Println("enter the num to search in linkedlist...")
	var num int
	fmt.Scanln(&num)

	fmt.Println(obj.searching_val(num))
}
