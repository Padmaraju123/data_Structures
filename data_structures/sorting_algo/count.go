package main

import "fmt"

func CountSort(sli []int) []int {
	k := FindMax(sli)
	n := len(sli)
	count := make([]int, k+1)
	b := make([]int, n+1)
	for i := 0; i < n; i++ {
		count[sli[i]]++
	}

	for i := 1; i < k+1; i++ {
		count[i] = count[i] + count[i-1]
	}

	for i := 0; i < n; i++ {
		b[count[sli[i]]] = sli[i]
		count[sli[i]] = count[sli[i]] - 1

	}
	return b

}

// find max number in array
func FindMax(sli []int) int {
	max := sli[0]
	for _, v := range sli {
		max = v
	}
	return max
}

func main() {
	slice := []int{0, 2, 5, 4, 2, 1, 5, 0, 3, 4, 3, 6, 4, 8, 9, 20}
	fmt.Println(CountSort(slice))
}
