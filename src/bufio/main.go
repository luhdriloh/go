package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		line := scan.Text()
		words := strings.Split(line, " ")
		fmt.Printf("[")
		for _, val := range words {
			fmt.Printf("%s, ", val)
		}
		fmt.Printf("]\n\n")
	}

}
