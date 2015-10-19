package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

func ReturnType(data interface{}) (reflect.Type, error) {
	knownType := true

	switch data.(type) {
	case string:
		data = data.(string)
	case int:
		data = data.(int)
	case int32:
		data = data.(int32)
	case int64:
		data = data.(int64)
	case float32:
		data = data.(float32)
	case float64:
		data = data.(float64)
	default:
		knownType = false
	}

	dataType := reflect.TypeOf(data)
	fmt.Println(dataType)

	if !knownType {
		err := errors.New("not a known type")
		return dataType, err
	} else {
		return dataType, nil
	}
}

func main() {
	var stringy interface{}

	stringy = 80085

	stringyType, err := ReturnType(stringy)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reflect.TypeOf(stringy))
	fmt.Println(stringyType)

	x := stringy.(int)
	if x > 23 {
		fmt.Println("hi")
	}
}
