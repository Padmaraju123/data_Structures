/*Given an array arr[] of N elements, the task is to write a function to search a given element x in arr[]


Input: arr[] = {10, 20, 80, 30, 60, 50,110, 100, 130, 170}, x = 110;
Output: 6
Explanation: Element x is present at index 6

Input: arr[] = {10, 20, 80, 30, 60, 50,110, 100, 130, 170}, x = 175;
Output: -1
Explanation: Element x is not present in arr[].

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checking(val int, sli []string, size int) int {
	kk, _ := strconv.Atoi(sli[size-1])

	if kk == val {
		return size-1

	}
	
	return checking(val,sli,size-1)

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	sli_c, _, _ := reader.ReadLine()
	sent := string(sli_c)
	split_c := strings.Split(sent, " ")

	fmt.Println("enter the value to serch....")

	var val int
	fmt.Scanln(&val)

	le := len(split_c)

	fmt.Printf("the index position of given val in slice is %d", checking(val, split_c, le))
}
