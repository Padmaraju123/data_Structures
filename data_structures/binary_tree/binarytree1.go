package main

import "fmt"

//create a node class
type Node struct {
	left    *Node
	right   *Node
	element int
}

//create a Tree class
type Tree struct {
	root *Node
}

func (obj *Tree) addlast(val int) {

	if obj.root == nil {
		obj.root = &Node{element: val, left: nil, right: nil}
	} else {
		obj.root.insert(val)
	}
}

func (n *Node) insert(val int) {

	if val < n.element {
		if n.left == nil {
			n.left = &Node{element: val, left: nil, right: nil}
		} else {
			n.left.insert(val)
		}
	} else {
		if n.right == nil {
			n.right = &Node{element: val, left: nil, right: nil}
		} else {
			n.right.insert(val)
		}
	}
}

//preorder traversing
func (obj *Tree) preorder(troot *Node) {
	if troot != nil {
		fmt.Print(troot.element, "-->")
		obj.preorder(troot.left)
		obj.preorder(troot.right)
	}
}

//inorder traversing
func (obj *Tree) inorder(troot *Node) {
	if troot != nil {
		obj.inorder(troot.left)
		fmt.Print(troot.element, "-->")
		obj.inorder(troot.right)
	}
}

//post order traversing
func (obj *Tree) postorder(troot *Node) {
	if troot != nil {
		obj.postorder(troot.left)
		obj.postorder(troot.right)
		fmt.Print(troot.element, "-->")
	}
}

//count of binary tree
func (obj *Tree) count(troot *Node) int {
	if troot != nil {
		x := obj.count(troot.left)
		y := obj.count(troot.right)
		return x + y + 1
	}
	return 0
}

//height of the tree
func (obj *Tree) Height(troot *Node) int {
	if troot != nil {
		x := obj.Height(troot.left)
		y := obj.Height(troot.right)
		if x > y {
			return x + 1
		} else {
			return y + 1
		}
	}
	return -1
}

//searching a val in binary tree
func (obj *Tree)searching(val int)bool{
	troot := obj.root
	for troot!=nil{
		if troot.element==val{
			return true
		}else if troot.element>val{
			troot = troot.left
		}else{
			troot = troot.right
		}
	}
	return false
	
}

//searching a val by recursion in binary tree
func (obj *Tree)searching_rec(troot *Node,val int)bool{
	if troot!=nil{
		if troot.element == val{
			return true
		}else if val < troot.element{
			return obj.searching_rec(troot.left,val)
		}else{
			return obj.searching_rec(troot.right,val)
		}
	}
	return false
}

//insertion val in the binary tree
var temp *Node
func (obj *Tree)insertion(val int){
	troot := obj.root
	newnode := Node{element: val,left:nil,right:nil}
	if val == troot.element{
		fmt.Println("same values in binary search tree is not applicable")
	}else{
		for troot!=nil{
			temp = troot
			if val < troot.element{
				troot = troot.left
			}else{
				troot = troot.right
			}
		}
	}
	
	if val < temp.element{
		temp.left = &newnode
	}else{
		temp.right = &newnode
	}

}

func main() {
	obj := Tree{}
	fmt.Print("enter the size of sli\n")
	var size int
	fmt.Scanln(&size)
	sli := make([]int, size)

	fmt.Print("enter the number line with one space\n")
	//25 15 30 10 20 35 28
	for i := 0; i < size; i++ {
		fmt.Scan(&sli[i])
	}

	for _, v := range sli {
		obj.addlast(v)
	}

	obj.preorder(obj.root)
	fmt.Println()

	obj.inorder(obj.root)
	fmt.Println()

	obj.postorder(obj.root)
	fmt.Println()

	fmt.Printf("the count of binary tree is %d\n", obj.count(obj.root))

	fmt.Printf("the height of the binary tree is %d\n", obj.Height(obj.root))

	fmt.Printf("the status of key is %v\n",obj.searching(30))

	fmt.Printf("the status of key is %v\n",obj.searching_rec(obj.root,30))

	obj.insertion(5)

	obj.preorder(obj.root)
}
