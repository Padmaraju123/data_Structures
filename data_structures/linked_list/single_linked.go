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

//adding nodes to list
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

//printing values and searching as well
func (obj *List) traversing_searching(val int) {
	comm := obj.head
	status := false

	for comm != nil {
		if comm.element == val {
			status = true
		}
		fmt.Print(comm.element, "-->\t")
		comm = comm.next
	}

	if status {
		fmt.Println("we found the value in slice")
	} else {
		fmt.Println("we didn't find the value in slice")
	}
}

//prepending the node in the list
func (obj *List) prepending(val int) {
	newnode := Node{element: val}

	if obj.len == 0 {
		obj.head, obj.tail = &newnode, &newnode
	} else {
		newnode.next, obj.head = obj.head, &newnode
		obj.len++
	}
}

//deletion the prepend value in list
func (obj *List) deletion_prepend() {

	if obj.len == 0 {
		fmt.Println("no nodes are occur")
		return
	}
	obj.head = obj.head.next
	obj.len--
}

//insertion in between nodes
func (obj *List)insertionBW(pos,val int){
	newnode := Node{element: val}

	if obj.len==0{
		obj.head,obj.tail = &newnode,&newnode
		return
	}

	comm := obj.head
	for i:=0;i<pos-1;i++{
		comm = comm.next
	}

	newnode.next,comm.next =comm.next,&newnode
	obj.len++
}

func (obj *List) Removebw(value int) {

	if obj.len == 0 {
		fmt.Println("no nodes are found")
		return 
	}

	if obj.head.element == value {
		obj.head = obj.head.next
		if obj.head == nil {
			obj.tail = nil
		}
		return
	}

	current := obj.head
	for current.next != nil {
		if current.next.element == value {
			current.next = current.next.next
			if current.next == nil {
				obj.tail = current
			}
			return
		}
		current = current.next
	}

}

//sorting the linked list

// Swap the given node value
func (this *List) swapData(a, b *Node) {
    temp := a.element;
    a.element = b.element;
    b.element = temp;
}

// Sort the linked list using selection sort
func (obj *List) selectionSort() {

    var  auxiliary *Node = obj.head;
    var  temp *Node = nil;
    var  node *Node = nil;
   
    for(auxiliary != nil) {

        node = auxiliary;
        temp = auxiliary.next;

        for(temp != nil) {
            if (node.element > temp.element) {
                node = temp;
            }
            temp = temp.next;
        }

        if (auxiliary.element > node.element) {
            obj.swapData(auxiliary, node);
        }
        auxiliary = auxiliary.next;
    }
}

func (obj *List)deletionend(){
	if obj.len==0{
		fmt.Println("no nodes found")
		return
	}
	cm := obj.head   
	for i:=0;i<obj.len-2;i++{
		cm=cm.next

	}
	cm.next=
	obj.len--
}

func main() {
	obj := List{}
	var size, search_value,prepend_value,pos,val,remove_val int
	fmt.Scanln(&size)

	sli := make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Scan(&sli[i])
	}

	for _, v := range sli {
		obj.appending(v)
	}

	fmt.Println("enter the value to search")
	fmt.Scan(&search_value)

	obj.traversing_searching(search_value)

	fmt.Println("enter the prepend value")
	fmt.Scan(&prepend_value)

	obj.prepending(prepend_value)

	obj.traversing_searching(search_value)

	obj.deletion_prepend()

	obj.traversing_searching(search_value)

	fmt.Println("enter postion and value to insertion in between")
	fmt.Scan(&pos,&val)

	obj.insertionBW(pos,val)

	obj.traversing_searching(search_value)

	fmt.Println("enter the value to delete")
	fmt.Scan(&remove_val)

	obj.Removebw(remove_val)

	obj.traversing_searching(search_value)

	obj.selectionSort()
	obj.traversing_searching(search_value)

	obj.deletionend()
	obj.traversing_searching(search_value)

}
