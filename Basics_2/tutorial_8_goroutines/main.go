package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

// wait group keeps track of requests that the CPU needs to switch back to after they return
var wg = sync.WaitGroup{}

var dbQueries = []string{"id1", "id2", "id3", "id4", "id5"}

func main() {
	t_begin := time.Now()

	// without go routines
	for i:=0; i < len(dbQueries); i++ {
		dbCall1(i)
	}
	fmt.Printf("Total execution time without goroutines: %v\n\n", time.Since(t_begin))

	t_begin = time.Now()
	// with go routines
	for i:=0; i < len(dbQueries); i++ {
		wg.Add(1) // add req to wait group
		go dbCall2(i) // now this will work concurrently
	}

	// wait for all requests to return
	wg.Wait()

	fmt.Printf("Total execution time with goroutines: %v\n\n", time.Since(t_begin))
}

func dbCall1(i int) {
	var delay float32 = rand.Float32()*2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("Query response: ", dbQueries[i])
}

func dbCall2(i int) {
	var delay float32 = rand.Float32()*2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("Query response: ", dbQueries[i])
	wg.Done()
}


