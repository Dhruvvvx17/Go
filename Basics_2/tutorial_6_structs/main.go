package main

import "fmt"

type gasEngine struct {
	mpg uint8
	gallons uint8
}
// method tied to the struct
func (e gasEngine) milesLeft() uint8 {
	return e.mpg * e.gallons
}

func main() {

	var myEngine1 gasEngine = gasEngine{mpg: 25, gallons: 15}
	fmt.Printf("Mileage: %v and Gallons remaining: %v\n", myEngine1.mpg, myEngine1.gallons)

	var myEngine2 gasEngine = gasEngine{30, 12}
	fmt.Printf("Mileage: %v and Gallons remaining: %v\n", myEngine2.mpg, myEngine2.gallons)

	var myEngine3 gasEngine
	myEngine3.mpg = 16
	myEngine3.gallons = 40
	fmt.Printf("Mileage: %v and Gallons remaining: %v\n", myEngine3.mpg, myEngine3.gallons)

	fmt.Println("Miles left in engine1: ", myEngine1.milesLeft())
	fmt.Println("Miles left in engine2: ", myEngine2.milesLeft())
	fmt.Println("Miles left in engine3: ", myEngine3.milesLeft())
}