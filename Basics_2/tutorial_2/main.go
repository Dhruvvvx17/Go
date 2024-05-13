// calculator
package main

import (
	"fmt"
	"errors"
)

func main() {
	var numerator int = 10
	var denominator int = 2

	var result, remainder, err = divide(numerator, denominator)

	switch {
		case err != nil:
			fmt.Printf(err.Error())
		case remainder == 0:
			fmt.Printf("The result of the integer division is %v", result)
		default:
			fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	}
}

func divide(numerator int, denominator int) (int, int, error) {

	var err error
	if denominator == 0 {
		err = errors.New("Cannot divide as denominator is zero")
		return 0, 0, err
	}


	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}