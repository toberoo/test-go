package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// Count will create a bunch of go
// routines to update the value
func Count() {
	var a int64

	for i := 0; i < 100; i++ {
		go func() {
			// Infinite loops
			for {
				atomic.AddInt64(&a, 1)
				time.Sleep(10 * time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("We counted: ", atomic.LoadInt64(&a), " times")
}
