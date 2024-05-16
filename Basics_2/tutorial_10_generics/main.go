package main

import "fmt"

func main() {
	var slice1 = []int{1,2,3,4,5}
	var slice2 = []float32{10.0, 2.0, 3.5, 4.2, 5.0}

	fmt.Printf("Sum of integer slice = %v\n", sumSlice[int](slice1))
	fmt.Printf("Sum of float slice = %v\n", sumSlice[float32](slice2))
}

func sumSlice[T int | float32](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}