package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			// Expression to be sent can be any suitable value.
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}

func fanIn(msg1, msg2 string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- msg1
			time.Sleep(time.Duration(rand.Intn(1e2)) * time.Millisecond)
		}
	}()
	go func() {
		for {
			c <- msg2
			time.Sleep(time.Duration(rand.Intn(1e2)) * time.Millisecond)
		}
	}()

	return c
}

func main() {
	whoWon := fanIn("levi", "sam")

	time.Sleep(time.Duration(time.Second * 2))

	for i := 0; i < 10; i++ {
		// Receive expression is just a value.
		fmt.Printf("I win: %q\n", <-whoWon)
	}

	fmt.Println("You're boring; I'm leaving.")
}
