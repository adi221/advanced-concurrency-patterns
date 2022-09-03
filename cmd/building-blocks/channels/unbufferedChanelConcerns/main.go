package main

import (
	"bytes"
	"fmt"
	"os"
)

/*
	if a goroutine making writes to a channel has knowledge of how many writes it will make,
	it can be useful to create a buffered channel whose capacity is
	the number of writes to be made, and then make those writes as quickly as possible.
*/
func main() {
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)

	intChan := make(chan int, 10)
	go func() {
		defer close(intChan)
		defer fmt.Println(&stdoutBuff, "Producer Done.")
		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
			intChan <- i
		}
	}()

	for {
		integer, ok := <-intChan
		fmt.Printf("(%v) Received %v.\n", ok, integer)
		if !ok {
			break
		}
	}
}
