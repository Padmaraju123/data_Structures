package main

import "fmt"


func reversing(sli []int)[]int{
	var rev_arr []int
	le := len(sli)-1
	for le>=0{
		rev_arr=append(rev_arr,sli[le])
		le--
	}
	return rev_arr
}


func main(){
	fmt.Println("enter the size of the array....")
	var size int
	fmt.Scanln(&size)

	arr := make([]int,size)

	fmt.Println("enter the number line with spaces")
	for i:=0;i<size;i++{
		fmt.Scan(&arr[i])
	}

	fmt.Println(reversing(arr))
}