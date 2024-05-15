package main

import "fmt"

// GAS ENGINE STRUCT
type gasEngine struct {
	mpg uint8
	gallons uint8
	engineType string
}
// method tied to the struct
func (e gasEngine) milesLeft() uint8 {
	return e.mpg * e.gallons
}

// ELECTRIC ENGINE STRUCT
type electricEngine struct {
	mpkWh uint8
	kWh uint8
	engineType string
}
// method tied to the struct
func (e electricEngine) milesLeft() uint8 {
	return e.mpkWh * e.kWh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, milesTarget uint8) {
	if milesTarget < e.milesLeft() {
		fmt.Printf("%v miles can be achieved\n", milesTarget)
	} else {
		fmt.Printf("%v miles cannot be achieved\n", milesTarget)
	}
}

func main() {

	var myEngine1 gasEngine = gasEngine{mpg: 25, gallons: 15, engineType: "gas"}
	fmt.Printf("Mileage: %v and Gallons remaining: %v\n", myEngine1.mpg, myEngine1.gallons)

	var myEngine2 gasEngine = gasEngine{30, 12, "gas"}
	fmt.Printf("Mileage: %v and Gallons remaining: %v\n", myEngine2.mpg, myEngine2.gallons)

	var myEngine3 gasEngine
	myEngine3.mpg = 16
	myEngine3.gallons = 40
	fmt.Printf("Mileage: %v and Gallons remaining: %v\n", myEngine3.mpg, myEngine3.gallons)

	fmt.Printf("\nEngine 1 [%s]: \n", myEngine1.engineType)
	canMakeIt(myEngine1, 50)
	canMakeIt(myEngine1, 126)


	var myEngine4 electricEngine
	myEngine4.mpkWh = 10
	myEngine4.kWh = 13
	myEngine4.engineType = "electric"

	fmt.Printf("\nEngine 4 [%s]: \n", myEngine4.engineType)
	canMakeIt(myEngine4, 50)
	canMakeIt(myEngine4, 150)
}