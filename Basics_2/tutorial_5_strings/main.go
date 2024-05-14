// string and runes

package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "resume"
	fmt.Println(myString)

	indexed := myString[0]

	fmt.Printf("%v, %T\n", indexed, indexed)

	for i, v := range(myString) {
		fmt.Println(i, v)
	}

	fmt.Printf("The number of bytes in the string: %v", len(myString))

	var myRune = 'a'
	fmt.Printf("\nmyRune = %v\n", myRune)

	// string concatenation
	// slice of characters (runes)
	var strSlice = []string{"d", "h", "r", "u", "v"}
	var cat = ""
	for i := range(strSlice) {
		cat += strSlice[i]
		fmt.Println(i, strSlice[i])
	}
	fmt.Printf("\nConcatenated string: %v\n", cat)


	var builder strings.Builder
	for i := range(strSlice) {
		builder.WriteString(strSlice[i])
	}
	var res = builder.String()

	fmt.Printf("\nConcatenated string using strings builder: %v\n", res)
}