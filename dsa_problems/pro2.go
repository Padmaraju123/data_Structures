//3.Given an array of integers, our task is to write a program that efficiently 
//finds the second-largest element present in the array.

package main

import  "fmt"

func second_big(sli []int)int{
	c := 2
	k := sli[0]
	for _,v := range sli[1:]{
		if v > k && c>1{
			k = v
			c--
		}
	}
	return k
}

func main(){
	var size int

	fmt.Scanln(&size)

	sli := make([]int,size)

	var i int

	for i<size{
		fmt.Scan(&sli[i])
		i++
	}

	fmt.Println(second_big(sli))
}