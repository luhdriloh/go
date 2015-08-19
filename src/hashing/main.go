package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func Salt() ([]byte, error) {
	numBytes := 8
	randBytes := make([]byte, numBytes)
	_, err := rand.Read(randBytes)

	return randBytes, err
}

func main() {
	password := []byte("monkeys $uck things")

	salt, err := Salt()
	if err != nil {
		fmt.Println(err)
	}

	shaHasher := sha256.New()

	for i := 0; i < 989; i++ {
		shaHasher.Write(append(salt, password...))
		password = shaHasher.Sum(nil)
	}

	fmt.Println(shaHasher.Sum(nil))
}
