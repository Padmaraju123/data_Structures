package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*func binary_search(sli []int, sk int) {
	le := len(sli)
	if le%2 != 0 {
		md := le / 2
		if sli[md] == sk {
			return md
		}
		if sk < sli[md] {
			binary_search(sli[:md], sk)
		}
	} else {
		md = le / 2

	}
}*/

func selection_ssort(sli []int) []int {
	le := len(sli)
	for j,_ := range sli[:le]{
		pos := j
		for i:=j+1;i<le;i++ {
			if sli[i] < sli[pos] {
				pos = i
			}
		}
		temp := sli[j]
		sli[j] = sli[pos]
		sli[pos] = temp

	}

	return sli
}

func main() {
	fmt.Println("enter the number line with spaces....")
	reader := bufio.NewReader(os.Stdin)
	sli_c, _, _ := reader.ReadLine()
	sent := string(sli_c)
	split_c := strings.Split(sent, " ")

	int_c := []int{}
	for _, v := range split_c {
		kk, _ := strconv.Atoi(v)
		int_c = append(int_c, kk)
	}

	fmt.Println("enter the value to search...")
	var sk int
	fmt.Scanln(&sk)
	//binary_search(int_c,sk)
	fmt.Println(selection_ssort(int_c))
}
