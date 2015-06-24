package main

import "fmt"

func FindDuplicate(search_list []int) []int {

	doubles_list := make([]int, 0)

	for _, value := range search_list {
		if InSlice(doubles_list, value) {
			doubles_list = append(doubles_list, value)
		}
	}

	return doubles_list
}

func InSlice(search_list []int, value_to_search int) bool {
	for _, value := range search_list {
		if value_to_search == value {
			return false
		}
	}

	return true
}

func main() {
	duplicate_list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 2, 4, 6, 8, 10, 2, 4}
	fmt.Println(duplicate_list)

	no_duplicate_list := FindDuplicate(duplicate_list)
	fmt.Println(no_duplicate_list)
}
