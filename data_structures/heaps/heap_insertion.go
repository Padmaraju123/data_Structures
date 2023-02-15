package main

import "fmt"

//creating a heap_array class

type heap struct {
	count int
	sli   []int
	csize int
}

func (obj *heap) creating() []int {
	obj.sli = make([]int, obj.count)
	return obj.sli
}

//length method to return length of the slice
func (obj *heap) length() int {
	return len(obj.sli)
}

//function which returns if length is zero
func (obj *heap) isempty() bool {
	if obj.length() == 0 {
		return true
	}
	return false
}

//insertion method to insert into binary tree
func (obj *heap) insertion(val int) {
	if obj.csize == obj.count {
		fmt.Println("No space is available to do")
		return
	}

	obj.csize = obj.csize + 1
	h1 := obj.csize

	for h1 > 1 && (val > obj.sli[h1/2]) {
		obj.sli[h1] = obj.sli[h1/2]
		h1 = h1 / 2
	}
	obj.sli[h1] = val
}

func (obj *heap) max() int {
	if obj.csize == 0 {
		fmt.Println("heap is empty")
		return -1
	}
	return obj.sli[1]
}

//deletion a key in the heap
func (obj *heap) deletion() int {

	if obj.isempty() {
		fmt.Println("heap is empty..")
		return 0
	}

	e := obj.sli[1]
	obj.sli[1] = obj.sli[obj.csize]
	obj.sli[obj.csize] = 0
	obj.csize = obj.csize - 1

	i := 1
	j := i * 2

	for j <= obj.csize {
		if obj.sli[j] < obj.sli[j+1] {
			j = j + 1
		}
		if obj.sli[i] < obj.sli[j] {
			temp := obj.sli[i]
			obj.sli[i] = obj.sli[j]
			obj.sli[j] = temp

			i = j
			j = i * 2
		} else {
			break
		}
	}
	return e

}

func main() {
	obj := heap{}
	//creating a slice with count size
	fmt.Scanln(&obj.count)
	obj.creating()

	obj.insertion(25)
	obj.insertion(14)
	obj.insertion(2)
	obj.insertion(20)
	obj.insertion(10)

	fmt.Println(obj.sli)

	//lets insertion one element
	obj.insertion(100)

	fmt.Println(obj.sli)

	//deleting the maximum value in heap
	fmt.Println(obj.deletion())

	fmt.Println(obj.sli)

}
