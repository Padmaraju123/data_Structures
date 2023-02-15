package main

import "fmt"

type Node struct {
	left  *Node
	right *Node
	data  int
}

//inserting nodes in the tree
func (obj Node) inserting(val int) {

	if obj.data != 0 {

		if val < obj.data {

			if obj.left == nil {
				obj.left = &Node{data: val}
			} else {
				obj.left.inserting(val)
			}

		} else if val > obj.data {

			if obj.right == nil {
				obj.right = &Node{data: val}
			} else {
				obj.right.inserting(val)
			}
		}
	} else {
		obj.data = val
	}

}

func (obj *Tree) preorder(troot *Node) {
	if troot != nil {
		fmt.Printf("%d ", troot.data)
		obj.preorder(troot.left)
		obj.preorder(troot.right)
	}
}

func main() {
	obj := Node{data: 12}
	obj.inserting(6)
	obj.inserting(14)
	obj.inserting(3)
	obj.inserting(34)
	obj.inserting(100)

	fmt.Println(obj.preorder(obj.data))

}
