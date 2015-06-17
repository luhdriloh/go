package main

import "fmt"

func main() {
	number_slice := make([]int, 100)

	for i, slice := range number_slice {
		slice = i
		fmt.Printf("[%d]\n", slice)
	}

	fmt.Println()

	// now a 2-d slice

	rows := 50
	columns := 20

	two_d_slice := make([][]int, rows)

	for i := 0; i < rows; i++ {
		two_d_slice[i] = make([]int, columns)

		for j := 0; j < columns; j++ {
			two_d_slice[i][j] = i * j
			fmt.Printf("[%d]", two_d_slice[i][j])
		}
		fmt.Println()
	}
}
