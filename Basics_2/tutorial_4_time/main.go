package main

import (
	"fmt"
	"time"
)

func main() {
	var n int = 1000000
	var testSlice1 = []int{}
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("Total execution time without preallocation: %v\n", timeloop(testSlice1, n))
	fmt.Printf("Total execution time with preallocation: %v\n", timeloop(testSlice2, n))
}

func timeloop(slice []int, n int) time.Duration {
	t0 := time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}