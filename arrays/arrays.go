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

	a := []int{14, 662, 1233, 464, 1355}

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	fmt.Println()

	// now the same thing but using range

	for i, v := range a {
		fmt.Printf("index %d : %d\n", i, v)
	}

}
