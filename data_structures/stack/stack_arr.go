package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*type stack struct {
	_ []int
}

func (obj *stack) length() int {
	return len(obj._data)
}

func (obj *stack) isempty() bool {
	if obj.length() == 0 {
		return true
	}
	return false
}

func (obj *stack) pushing(val int) {
	obj._data = append(obj._data, val)
}

func (obj *stack) poping() {
	
	if obj.isempty() {
		fmt.Println("there is no elements to pop in stack....")

	} else {
		le := obj.length() - 1
		for le >= 0 {
			fmt.Printf("%d\n", obj._data[le])
			le--
		}
	}
}*/

type stack struct{
	sli []int
}

func (obj *stack)length()int{
	return len(obj.sli)
}


func (obj *stack)pushing(val int){
	obj.sli = append(obj.sli,val)
}

func (obj *stack)poping(){
	if obj.length()==0{
		fmt.Println("No elements are found in stack to pop....")
		return
	}
	le := obj.length()-1

	for le>=0{
		fmt.Print(obj.sli[le],"--->")
		le--
	}

}


func main() {

	obj := stack{}

	fmt.Print("enter the number line with spaces\n")

	reader := bufio.NewReader(os.Stdin)
	sli_c, _, _ := reader.ReadLine()
	sent := string(sli_c)

	split_c := strings.Split(sent, " ")

	for _, v := range split_c {
		vv, _ := strconv.Atoi(v)
		obj.pushing(vv)
	}

	fmt.Print("printing the values in stack in LILO...\n")
	obj.poping()

}
