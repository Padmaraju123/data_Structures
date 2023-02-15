/*Given an array with all distinct elements, find the largest three elements. Expected time complexity is O(n) and extra space is O(1).

Examples :

Input: arr[] = {10, 4, 3, 50, 23, 90}
Output: 90, 50, 23*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func working(sli []int) {
	sort.Ints(sli)
	le := len(sli)
	for _,w := range sli[le-3:]{
		fmt.Println(w)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	sli_c, _, _ := reader.ReadLine()

	sent := string(sli_c)

	split_c := strings.Split(sent, " ")

	int_S := []int{}

	for _, j := range split_c{
		kh,_ := strconv.Atoi(j)
		int_S = append(int_S, kh)
	}

	working(int_S)

}
