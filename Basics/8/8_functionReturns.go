package main

import "fmt"

func getName() (string, string) {
	return "Dhruv", "Vohra"
}

func main() {
	firstName, _ := getName()
	fmt.Printf("Hi %s", firstName)
}