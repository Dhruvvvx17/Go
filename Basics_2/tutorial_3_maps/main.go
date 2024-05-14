// maps
package main

import "fmt"

func main() {
	// declaration
	// var <name> map[data_type_for_key]data_type_for_value = make(<same as LHS>)
	var cache map[string]uint8 = make(map[string]uint8)
	fmt.Println(cache)

	// declaration and initialization
	var cache2 = map[string]uint8{"Adam": 23, "Sarah": 45, "Eve": 25}
	fmt.Println(cache2)

	// accessing data

	// existing key
	fmt.Printf("Adam's age: %v\n",cache2["Adam"])

	// non-existing key: returns default value according to value data type
	fmt.Printf("Jason's age: %v\n",cache2["Jason"])

	// access with check for existance
	var age, exists = cache2["Jason"]
	if exists {
		fmt.Printf("The age is %v\n", age)
	} else {
		fmt.Printf("Invalid Name\n")
	}

	// deletion
	delete(cache2, "Sarah")
	fmt.Println(cache2)

	// iterating the map - order is not preserved
	for name := range cache2 {
		fmt.Printf("Name: %v\n", name)
	}
}