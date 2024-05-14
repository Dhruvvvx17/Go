package main

import "fmt"

func main() {

	// while loop using 'for' keyword
	fmt.Printf("While loop behaviour: ")
	var i int = 0
	for i < 10 {
		fmt.Printf("%v ",i)
		i += 1
	}

	// while loop using 'for' keyword and break
	fmt.Printf("\nWhile loop behaviour with break: ")
	i = 0
	for {
		if i >= 10 {
			break
		}
		fmt.Printf("%v ",i)
		i += 1
	}

	fmt.Printf("\nFor loop behaviour: ")
	for i:=0; i<10; i++ {
		fmt.Printf("%v ",i)
	}
	
}