package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	sentence_line := "welcome to all"

	//creating a file and give as parameter as path where to create
	file, err := os.Create("./sent.txt")

	if err != nil {
		panic(err)
	}

	//writing a file it return length and error
	length, error := io.WriteString(file, sentence_line)

	if error != nil {
		panic(error)
	}

	fmt.Println(length)

	//how to close a file
	file.Close()

}
