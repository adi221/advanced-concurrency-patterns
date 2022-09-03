package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			default:
			}

			// Do non-preemptable work
			fmt.Println("First loop")
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			select {
			case <-done:
				return
			default:
				// Do non-preemptable work
				fmt.Println("Second loop")
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	done <- true
}
