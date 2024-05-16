package main

import (
	"fmt"
	"time"
	"sync"
)

// to secure lock a write operation
var m = sync.Mutex{}

// wait group keeps track of requests that the CPU needs to switch back to after they return
var wg = sync.WaitGroup{}

var dbQueries = []string{"id1", "id2", "id3", "id4", "id5"}
var dbResults = []string{}

func main() {
	t_begin := time.Now()

	// with go routines and locks
	for i:=0; i < len(dbQueries); i++ {
		wg.Add(1) // add req to wait group
		go dbCallConcurrent2(i) // now this will work concurrently
	}

	// wait for all requests to return
	wg.Wait()

	fmt.Printf("Final result: %v\n", dbResults)
	fmt.Printf("Total execution time with goroutines: %v\n\n", time.Since(t_begin))
}

func dbCallConcurrent2(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("Query response: ", dbQueries[i])

	// lock before writing
	m.Lock()
	// write
	dbResults = append(dbResults, dbQueries[i])
	// unlock after writing
	m.Unlock()

	wg.Done()
}


