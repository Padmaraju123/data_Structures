package main

import (
	"fmt",
)

func Insertionsort(items []int)[]int{
	n := len(items)
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
		}
	}
	return items
}

func main() {
	sli := []int{12, 34, 5, 6, 66}
	fmt.Println(Insertionsort(sli))
}

