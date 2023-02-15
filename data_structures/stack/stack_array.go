package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stack_Arr struct {
	sli []int
}

func (obj *stack_Arr)length()int{
	return len(obj.sli)
}

func (obj *stack_Arr)empty()int{
	if len(obj.sli)==0{
		return 0
	}
	return -1
}

func (obj *stack_Arr) pushing(val int) {
	obj.sli = append(obj.sli, val)
}

func (obj *stack_Arr) printing() {
	for _, v := range obj.sli {
		fmt.Printf("%T %v\n", v, v)
	}
}

func (obj *stack_Arr)poping(){
	if obj.length()==0{
		fmt.Println("no values have in stack")
		
	}else{
		le := obj.length()-1
		for le>=0{
			fmt.Println(obj.sli[le])
			le--
		}
	}
}

func main() {
	obj := stack_Arr{}
	reader := bufio.NewReader(os.Stdin)
	sli_b, _, _ := reader.ReadLine()
	sent := string(sli_b)

	split_S := strings.Split(sent, " ")

	for _, v := range split_S {
		vv, _ := strconv.Atoi(v)
		obj.pushing(vv)
	}

	fmt.Print("printing the values from latest to old\n")
	obj.poping()



}
