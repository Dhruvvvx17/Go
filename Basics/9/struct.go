package main

import "fmt"

type messageToSend struct {
	message string
	phoneNumber int
}

func test(m messageToSend) {
	fmt.Printf("Sending message %s to %v\n", m.message, m.phoneNumber)
}

func main() {

	test(messageToSend {
		message: "Hello Dhruv!",
		phoneNumber: 6028155497,
	})

	test(messageToSend {
		message: "Hello World!",
		phoneNumber: 1234567890,
	})
}