package main

import "fmt"

func main() {
	const word = "hello world"
	var word_array [len(word)]string

	// put the chars that are in var word into var word_array

	for i := 0; i < len(word); i++ {
		word_array[i] = string(word[i])
	}

	//  print out our letters using the array

	for i := 0; i < len(word); i++ {
		fmt.Println(word_array[i], ",\n")
	}

	// initialize an int array

	a := []int{14, 662, 1233, 464, 1355}

	// print out with a regular loop

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	fmt.Println()

	// now the same thing but using range

	for i, v := range a {
		fmt.Printf("index %d : %d\n", i, v)
	}

	fmt.Println()

	///////////////////////////////////////////
	//************** 2-d array **************//
	///////////////////////////////////////////

	var two_array [10][5]int

	for i := 0; i < 10; i++ {
		for j := 0; j < 5; j++ {
			two_array[i][j] = i * j
			fmt.Printf("[%d]", two_array[i][j])
		}
		fmt.Println()
	}

}
