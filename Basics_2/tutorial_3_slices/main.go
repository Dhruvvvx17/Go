// slices
package main

import "fmt"

func main() {
	// declaration
	// note: no length value provided
	var slice []int32 = []int32{10, 20, 30, 40}

	fmt.Println(slice)
	fmt.Printf("Address of slice: %v\n", &slice[0])
	fmt.Printf("The length of slice is %v with capacity %v\n", len(slice), cap(slice))

	// access - by index
	fmt.Println(slice[0])

	// append to slice - address changed if length == capacity
	slice = append(slice, 50)

	fmt.Println(slice)
	fmt.Printf("Address of slice: %v\n", &slice[0])
	fmt.Printf("The length of slice is %v with capacity %v\n", len(slice), cap(slice))

	// appending slice to a slice - using SPREAD ... operator
	var newSlice []int32 = []int32{9, 10, 11}
	slice = append(slice, newSlice...)

	fmt.Println(slice)
	fmt.Printf("Address of slice: %v\n", &slice[0])
	fmt.Printf("The length of slice is %v with capacity %v", len(slice), cap(slice))


	// iterating
	for idx, val := range newSlice {
		fmt.Printf("%v: %v\n", idx, val)
	}	
}