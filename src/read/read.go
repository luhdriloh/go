package main

import (
	"bufio"
	"fmt"
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
	text_reader := bufio.NewScanner(file)

	var lines []string
	for text_reader.Scan() {
		lines = append(lines, text_reader.Text())
	}

	new_path_name := "read-" + path
	to_text, _ := os.Create(new_path_name)
	defer to_text.Close()

	for _, str := range lines {
		fmt.Fprintf(to_text, str)
	}

}
