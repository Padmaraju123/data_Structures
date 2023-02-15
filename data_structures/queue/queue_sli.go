package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type queue struct {
	sli []int
}

func (obj *queue)length()int{
	return len(obj.sli)
}

func (obj *queue)isempty()bool{
	if obj.length()==0{
		return true
	}
	return false
}

func (obj *queue) Enquening(val int){
	obj.sli = append(obj.sli, val)
}

func (obj *queue)Dequening(){
	if obj.isempty(){
		fmt.Println("no elements to Deqeue in sli....")
		return
	}
	
	for i:=0;i<obj.length();i++{
		fmt.Print(obj.sli[i],"--->")
	}
}

func main() {
	fmt.Println("enter the number line with spaces...")
	reader := bufio.NewReader(os.Stdin)
	sli_c, _, _ := reader.ReadLine()
	sent := string(sli_c)

	split_c := strings.Split(sent, " ")

	obj := queue{}

	for _, v := range split_c {
		kk, _ := strconv.Atoi(v)
		obj.Enquening(kk)
	}

	obj.Dequening()

}
