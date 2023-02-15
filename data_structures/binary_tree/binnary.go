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
	left    *Node
	right   *Node
}

type Tree struct {
	root *Node
}

// appending the tree with nodes
func (obj *Tree) getting(val int) {

	if obj.root == nil {
		obj.root = &Node{element: val, left: nil, right: nil}
		return
	}
	obj.root.insertion(val)

}

func (n *Node) insertion(val int) {
	if val < n.element {
		if n.left == nil {
			n.left = &Node{element: val, left: nil, right: nil}
		} else {
			n.left.insertion(val)
		}
	} else {
		if n.right == nil {
			n.right = &Node{element: val, left: nil, right: nil}
		} else {
			n.right.insertion(val)
		}
	}

}

// preorder traversal in binary tree
func preorder(troot *Node) {
	if troot == nil{ 
		return
	}
	fmt.Print(troot.element,"-->")
	preorder(troot.left)
	preorder(troot.right)
}






func main() {
	fmt.Println("enter the number line with spaces....")
	//100 80 105 50 90 101 120
	reader := bufio.NewReader(os.Stdin)
	sli_c, _, _ := reader.ReadLine()
	snt := string(sli_c)
	split_c := strings.Split(snt, " ")

	obj := Tree{}
	for _, v := range split_c {
		kk, _ := strconv.Atoi(v)
		obj.getting(kk)
	}

	preorder(obj.root)
	//100-->80-->50-->90-->105-->101-->120-->



}
