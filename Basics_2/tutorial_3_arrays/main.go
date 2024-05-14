// arrays
package main

import "fmt"

func main() {
	// declaration
	var arr [3]int32

	// assignment
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30

	// access - by index
	fmt.Println(arr[0])

	// access - by slice
	fmt.Println(arr[1:3])

	// memory location
	fmt.Println(&arr[0])
	fmt.Println(&arr[1])
	fmt.Println(&arr[2])

	// ---------------

	// inline declaration and initialization
	var newArr [3]int32 = [3]int32{5, 15, 23}
	fmt.Println(newArr)

	// syntactic sugar
	newArr2 := [...]int32{9,10,11}
	fmt.Println(newArr2)


	// iterating
	for idx, val := range newArr2 {
		fmt.Printf("%v: %v\n", idx, val)
	}	
}