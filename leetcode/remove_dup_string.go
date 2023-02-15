package main

import "fmt"

func remove_dup(word string) string {

	le := len(word)

	if le < 1{
		return word
	}else{
		for i := 0; i < le-1; i++ {
			if word[i] == word[i+1] {
				word = word[:i] + word[i+2:]
				return remove_dup(word)
			}
		}

	}
	return ""

}

func main() {
	var word string
	fmt.Println("enter the word with repeated letters...")
	fmt.Scanln(&word)
	fmt.Println(remove_dup(word))
}
