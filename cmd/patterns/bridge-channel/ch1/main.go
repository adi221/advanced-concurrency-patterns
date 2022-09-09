package main

import (
	"fmt"

	"github.com/adi221/playground/advanced-concurrency-patterns/pkg/concepts"
)

func main() {
	genVals := func() <-chan <-chan interface{} {
		chanStream := make(chan (<-chan interface{}))
		go func() {
			defer close(chanStream)
			for i := 0; i < 10; i++ {
				stream := make(chan interface{}, 1)
				stream <- i
				close(stream)
				chanStream <- stream
			}
		}()
		return chanStream
	}
	done := make(chan bool)
	defer close(done)

	for v := range concepts.Bridge(nil, genVals()) {
		fmt.Printf("Val is %v\n", v)
	}
}
