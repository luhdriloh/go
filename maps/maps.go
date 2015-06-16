package main

import "fmt"

func main() {
	a_map := make(map[string]int)

	a_map["uno"] = 1
	a_map["dos"] = 2
	a_map["tres"] = 3

	delete(a_map, "dos") // this deletes a map and key

	a_map["tres"] = 100 // reassign "tres" to a new value

	// this checks to see if this key is inside the map if it is not then ok will be
	//  false and value will the zero value for our chosen value type

	value, ok := a_map["quatro"]
	fmt.Println("The value:", value, "Present?", ok)

	value, ok = a_map["tres"]
	fmt.Println("The value:", value, "Present?", ok)

	for _, current_map := range a_map {
		fmt.Println(current_map)
	}
}
