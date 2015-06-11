package main

import (
	"fmt"
	"reflect"
)

// the var below is for package level variable declarations GLOBAL
var (
	name   string
	course string
	grade  float64

	dog, dog_year = "Memo", .3 // easier to read if broken up
)

func main() {
	fmt.Println("name is", name, "and is type", reflect.TypeOf(name), "\ngrade is", grade,
		"and is type", reflect.TypeOf(grade))

	fmt.Println("my dogs name is", dog, "and he is", dog_year, "years old")

	acceleration := 9.81
	velocity := 50

	weird_number := int(acceleration) * velocity

	fmt.Println(weird_number)
}
