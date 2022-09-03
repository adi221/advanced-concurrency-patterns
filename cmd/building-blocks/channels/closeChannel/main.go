package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

func main() {
	durationToBlockRoutines := flag.Int("durationToBlockRoutines", 3, "duration to block routines in seconds")

	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Printf("Channel number %d starts doing the work..\n", i)
		}(i)
	}
	t := time.Duration(*durationToBlockRoutines)
	time.Sleep(t * time.Second)
	close(begin) // Close the channel unblocks the routines
	wg.Wait()
}
