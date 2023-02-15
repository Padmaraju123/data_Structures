package main

import (
	"fmt"
)

func selectionsort(items []int) []int {
	length := len(items)
	for i := 0; i < length; i++ {
		var min = i
		for j := i; j < length; j++ {
			if items[j] < items[min] {
				min = j
			}
		}
		items[i], items[min] = items[min], items[i]
	}
	return items
}

func main() {

	sli := []int{12, 34, 5, 6, 66}
	fmt.Println(selectionsort(sli))

}
