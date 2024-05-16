package main

import (
	"fmt"
 	"time"
)

func main() {
	var c1 = make(chan int)
	var c2 = make(chan int, 5)

	// this call - doesn't let the process function end until the sleep cycles are completed
	go process(c1)
	for i := range c1{
		fmt.Println(i)
		time.Sleep(time.Second*1)
	}


	// this call - uses a buffer channel of 5, 
	// so it allows the process function to end immediately while main continues with its sleep cycles
	go process(c2)
	for i := range c2{
		fmt.Println(i)
		time.Sleep(time.Second*1)
	}
}


func process(c chan int) {

	// defer - do this statement right before returning from the function
	defer close(c)

	for i:=0; i<5; i++ {
		c <- i
	}

	fmt.Println("Exiting process...")
}