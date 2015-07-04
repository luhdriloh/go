package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// create a ReadWriter pointer

	path := "test.txt"

	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// file implements Read so it is a Reader interface
	text_reader, _ := ioutil.ReadFile(path)

	new_path_name := "read-" + path

	ioutil.WriteFile(new_path_name, text_reader, 0777)

	source, _ := os.Open("test2.txt")
	destination, _ := os.Create("read-test2.txt")

	bytesRead, _ := io.Copy(destination, source)

	fmt.Println(bytesRead)
}
