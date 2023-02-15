package main

import "fmt"

func receiving(sli []int)[]int{
	le := len(sli)-1
	reversing(&sli,start,le)
	return sli
}

var temp,start int
func reversing(sli *[]int,start,end int){

	if start <= end{
		temp = (*sli)[start]
		(*sli)[start]= (*sli)[end]
		(*sli)[end] = temp
		reversing(sli,start+1,end-1)
	}
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

	fmt.Println(receiving(arr))
}