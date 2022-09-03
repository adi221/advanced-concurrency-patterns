package main

import (
	"fmt"
	"time"
)

func main() {
	amountOfWorkers := 100
	c1 := make(chan int, amountOfWorkers)
	c2 := make(chan int, amountOfWorkers)

	go func() {
		defer close(c1)
		for i := 0; i < amountOfWorkers; i++ {
			c1 <- i
			time.Sleep(time.Millisecond * 50)
		}
	}()
	go func() {
		// defer close(c2) // if channel is closed, the Go runtime will perform a pseudo-random uniform selection over the set of case statements.
		for i := 0; i < amountOfWorkers; i++ {
			c2 <- i
			time.Sleep(time.Millisecond * 2)
		}
	}()

	for i := 0; i < amountOfWorkers*2; i++ {
		select {
		case c1Msg, opened := <-c1:
			fmt.Println("Message from c1", c1Msg, opened)
		case c2Msg, opened := <-c2:
			fmt.Println("Message from c2", c2Msg, opened)
		}
	}
}
