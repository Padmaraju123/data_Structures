package main

import "fmt"

func Bobblesort(sli []int) []int {
	for i := 0; i < len(sli)-1; i++ {
		for j := 0; j < len(sli)-i-1; j++ {
			if sli[j] > sli[j+1] {
				sli[j], sli[j+1] = sli[j+1], sli[j]
			}
		}
	}
	return sli
}
func main() {
	slice := []int{12, 14, 56, 20, 2, 78, 9, 756, 22, 78}
	fmt.Println(Bobblesort(slice))
}
