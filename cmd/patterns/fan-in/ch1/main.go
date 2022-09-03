package main

import (
	"fmt"
	"sync"
)

/*
In a nutshell, fanning in involves creating the multiplexed channel
consumers will read from, and then spinning up one goroutine for each incoming channel,
and one goroutine to close the multiplexed channel when the incoming channels have all been closed.
Since we’re going to be creating a goroutine that is waiting on N other goroutines to complete,
it makes sense to create a sync.WaitGroup to coordinate things. The multiplex function also notifies the WaitGroup that it’s done.
*/

type T interface {
	int32 | string
}

func makeRChan[k T](work []k) <-chan k {

	inputChannel := make(chan k)
	go func() {

		defer close(inputChannel)
		for _, v := range work {
			inputChannel <- v
		}

	}()
	return inputChannel
}

func fanIn[k T](channels ...<-chan k) <-chan k {

	var wg sync.WaitGroup
	combinedOutputChannel := make(chan k)
	wg.Add(len(channels))

	for _, o := range channels {
		go func(c <-chan k) {
			for {

				value, ok := <-c

				if !ok {
					wg.Done()
					break
				}

				combinedOutputChannel <- value

			}
		}(o)
	}

	go func() {

		wg.Wait()
		close(combinedOutputChannel)

	}()

	return combinedOutputChannel
}

func main() {

	i1 := makeRChan([]int32{0, 2, 6, 8})
	i2 := makeRChan([]int32{1, 5, 19, 23})

	out := fanIn(i1, i2)

	for value := range out {
		fmt.Println("Value:", value)
	}

}
