package main

// https://www.ardanlabs.com/blog/2017/10/the-behavior-of-channels.html

import (
	"fmt"
)

func main() {
	strChan := make(chan string)
	go func() {
		strChan <- "Hello"
	}()
	salute, ok := <-strChan
	fmt.Printf("%v, %v\n", salute, ok)
	close(strChan)
	secondSalute, ok := <-strChan
	fmt.Printf("%v, %v\n", secondSalute, ok)
}
