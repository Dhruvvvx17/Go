package main

import "fmt"

func main() {
	var p *int32 = new(int32)
	var i int32 = 5

	fmt.Printf("Value of i = %v\n", i)
	fmt.Printf("Value at p = %v\n", *p)
	fmt.Printf("p points to = %p\n", p)

	p = &i
	*p = 10

	fmt.Printf("Value of i = %v\n", i)
	fmt.Printf("Value at p = %v\n", *p)
	fmt.Printf("p points to = %p\n\n", p)

	// -----------

	var arr1 = [5]int32{1,2,3,4,5}

	fmt.Printf("memory location of arr1 is = %p\n", &arr1)
	var result1 = square_without_pointers(arr1)
	fmt.Printf("memory location of result1 is = %p\n", &result1)
	fmt.Printf("result1 elemets = %v\n", result1)
	fmt.Printf("array elemets = %v\n\n", arr1)

	fmt.Printf("memory location of arr1 is = %p\n", &arr1)
	var result2 = square_with_pointers(&arr1)
	fmt.Printf("memory location of result2 is = %p\n", &result2)
	fmt.Printf("result2 elemets = %v\n", result2)
	fmt.Printf("array elemets = %v\n\n", arr1)
}

func square_without_pointers(temp1 [5]int32) [5]int32 {
	fmt.Printf("memory location of temp1 is = %p\n", &temp1)
	for i := range(temp1) {
		temp1[i] = temp1[i] * temp1[i]
	}
	return temp1
}

func square_with_pointers(temp2 *[5]int32) [5]int32 {
	fmt.Printf("memory location of temp2 is = %p\n", temp2)
	for i := range(temp2) {
		temp2[i] = temp2[i] * temp2[i]
	}
	return *temp2
}