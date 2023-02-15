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

func (obj *Tree) appending(val int) {

	if obj.root == nil {
		obj.root = &Node{element: val, left: nil, right: nil}
		return
	}

	obj.root.insert(val)

}

func (n *Node) insert(val int) {

	if val <= n.element {

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

func (obj *Tree) preorder(troot *Node) {
	if troot != nil {
		fmt.Printf("%d ", troot.element)
		obj.preorder(troot.left)
		obj.preorder(troot.right)
	}

}

func (obj *Tree) Inorder(troot *Node) {
	if troot != nil {
		obj.Inorder(troot.left)
		fmt.Printf("%d ", troot.element)
		obj.Inorder(troot.right)
	}

}

func (obj *Tree) Postorder(troot *Node) {
	if troot != nil {
		obj.Postorder(troot.left)
		obj.Postorder(troot.right)
		fmt.Printf("%d ", troot.element)
	}

}

// count of nodes in binary tree
func (obj *Tree) counting(troot *Node) int {

	if troot != nil {
		x := obj.counting(troot.left)
		y := obj.counting(troot.right)
		return x + y + 1
	}
	return 0

}

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

// searching iterative method
func (obj *Tree) Searching(val int) bool {

	troot := obj.root
	for troot != nil {
		if troot.element == val {
			return true
		} else if troot.element < val {
			troot = troot.right
		} else {
			troot = troot.left
		}
	}
	return false
}

// searching recursive method
func (obj *Tree) searching_rec(troot *Node, key int) bool {
	if troot != nil {
		if troot.element == key {
			return true
		} else if key < troot.element {
			return obj.searching_rec(troot.left, key)
		} else {
			return obj.searching_rec(troot.right, key)
		}
	}
	return false
}

// insertion iterative method
var temp *Node

func (obj *Tree) insertion_val(val int) {

	troot := obj.root

	for troot != nil {
		temp = troot
		if val == troot.element {
			return
		} else if val > troot.element {
			troot = troot.right
		} else {
			troot = troot.left
		}
	}

	newnode := Node{element: val, left: nil, right: nil}

	if temp != nil {
		if val < temp.element {
			temp.left = &newnode
		} else {
			temp.right = &newnode
		}
	} else {
		troot = &newnode
	}
}

// insertion with recursion
func (obj *Tree) insertion_rec(troot *Node, val int) *Node {

	if troot != nil {
		if val < troot.element {
			troot.left = obj.insertion_rec(troot.left, val)
		} else {
			troot.right = obj.insertion_rec(troot.right, val)
		}
	} else {
		newnode := Node{element: val, left: nil, right: nil}
		troot = &newnode
	}
	return troot
}

func (obj *Tree) con_int(sli []string) {
	for _, v := range sli {
		vv, _ := strconv.Atoi(v)
		obj.appending(vv)
	}
	obj.preorder(obj.root)
	fmt.Println()

	obj.Inorder(obj.root)
	fmt.Println()

	obj.Postorder(obj.root)
	fmt.Println()

	fmt.Println(obj.counting(obj.root))

	fmt.Println(obj.Height(obj.root))
	fmt.Println(obj.Searching(10))

	fmt.Println(obj.searching_rec(obj.root, 10))

	obj.insertion_val(35)

	obj.preorder(obj.root)
	fmt.Println()
	add := obj.insertion_rec(obj.root, 72)

	obj.preorder(add)

}

func main() {
	obj := Tree{}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter numberline with spaces\n")
	// 5 3 8 1 4 6 9
	sli_b, _, _ := reader.ReadLine()
	sent := string(sli_b)
	split_s := strings.Split(sent, " ")
	obj.con_int(split_s)

}
