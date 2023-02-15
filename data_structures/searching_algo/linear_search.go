package main

import "fmt"

func linear_Searching(sli []int,val int)string{

	for _,v := range sli{
		if v==val{
			return "we found the value"
		}
	}
	return "we didn't found the value"

}


func main(){
	var size,val int
	fmt.Scan(&size,&val)

	sli := make([]int,size)
	fmt.Println(sli)
	for i:=0;i<size;i++{
		fmt.Scan(&sli[i])
	}
	
	fmt.Println(linear_Searching(sli,val))

}