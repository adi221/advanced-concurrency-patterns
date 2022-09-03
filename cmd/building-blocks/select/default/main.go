package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	var c1, c2 <-chan int
	select {
	case <-c1:
	case <-c2:
	// default statement runs almost instantaneously. This allows you to exit a select block without blocking.
	// Usually you’ll see a default clause used in conjunction with a for-select loop.
	default:
		fmt.Printf("In default after %v\n\n", time.Since(start))
	}
}
