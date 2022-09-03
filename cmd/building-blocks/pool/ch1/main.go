package main

import (
	"fmt"
	"sync"
)

func main() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance.")
			return &struct{}{}
		},
	}

	// Only one instance will be created
	instance1 := myPool.Get()
	myPool.Put(instance1)
	instance2 := myPool.Get()
	myPool.Put(instance2)
	instance3 := myPool.Get()
	myPool.Put(instance3)
}
