package main

import (
	"fmt"
	"math/rand"
	"time"
)

func add_values(val []int, start, end int) {
	for i := start; i < end; i++ {
		val[i] *= 14
	}
}

func main() {
	r := rand.New(rand.NewSource(79))
	vals := make([]int, 0)
	for i := 0; i < 40000; i++ {
		vals = append(vals, r.Int()%300)
	}

	newTimer := time.Now()
	InsertionSort(vals, 0, len(vals)-1)
	timeElapsed := time.Now()
	fmt.Printf("The call took %v to run.\n", timeElapsed.Sub(newTimer))

	good := true
	for i := range vals {
		if i == 0 {
			continue
		}
		if vals[i] < vals[i-1] {
			good = false
			break
		}
	}
	fmt.Println(good)
}
