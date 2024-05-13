package main

import "fmt"

func main() {
	
	// message := "Hello World!"
	messageLen := 12
	maxMessageLen := 20

	if messageLen <= maxMessageLen {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}
}