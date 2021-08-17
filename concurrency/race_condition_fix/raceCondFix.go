package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var c int32
var waitGroup sync.WaitGroup

// feel free to test doodle
func main() {
	// add waitgroup for each goroutine
	waitGroup.Add(3)

	go increment("yamaha")
	go increment("mt")
	go increment("25")

	// waiting for goroutines to complete
	waitGroup.Wait()

	fmt.Println(c)
}

func increment(some string) {
	// Done() func used to tell that is done
	defer waitGroup.Done()

	for range some {
		//Atomic Function
		//For fix race condition
		atomic.AddInt32(&c, 1)

		//enter thread in the line by line
		runtime.Gosched()
	}
}
