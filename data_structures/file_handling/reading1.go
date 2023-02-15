package main

import (
	"fmt"
	"io/ioutil"
)

func main(){
	sli_b,err :=ioutil.ReadFile("data.txt")
	_ = err
	fmt.Println(string(sli_b))
}