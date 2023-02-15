package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var val int

func binary_search(sl []int)int{
	fmt.Scanln(&val)
	r := len(sl)
	l := 0
	for l<=r{
		m := (l+r)/2
		if val==sl[m]{
			return m
		}else if val < sl[m]{
			r = m-1
		}else{
			l = m+1
		}
	}
	return -1

}

/*func selection_Sort(sli []int)int{
	le := len(sli)
	for i,_ := range sli[:le]{
		pos:=i
		for j:=i+1;j<le;j++{
			if sli[j]<sli[pos]{
				pos=j
			}
		}
		sli[i],sli[pos] = sli[pos],sli[i]
	
	}
	
	return binary_search(sli)

}*/

func selecti(sli []int)int{
	le := len(sli)
	for i,_ := range sli[:le-1]{
		pos := i
		for j:=i+1;j<le;j++{
			if sli[j]<sli[pos]{
				pos = j
			}
		}
		temp := sli[i]
		sli[i] = sli[pos]
		sli[pos] = temp
	}
	return binary_search(sli)
}

func main() {
	fmt.Println("enter the number line with spaces....")
	reader := bufio.NewReader(os.Stdin)
	sli_b, _, _ := reader.ReadLine()
	sent := string(sli_b)
	split_s := strings.Split(sent, " ")
	int_s := []int{}

	for _, v := range split_s {
		vv, _ := strconv.Atoi(v)
		int_s = append(int_s, vv)
	}

	//fmt.Println(selection_Sort(int_s))
	fmt.Println(selecti(int_s))
}
