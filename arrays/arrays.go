package main

import "fmt"

func main() {
	const word = "hello world"
	var word_array [len(word)]string

	for i := 0; i < len(word); i++ {
		word_array[i] = string(word[i])
	}

	for i := 0; i < len(word); i++ {
		fmt.Println(word_array[i], ",\n")
	}

}
