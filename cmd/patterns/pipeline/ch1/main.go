package main

import "fmt"

func main() {
	ints := []int{1, 2, 3, 4}
	for _, v := range add(multiply(ints, 2), 12) {
		fmt.Println(v)
	}
}

func multiply(values []int, multiplier int) []int {
	multipliedValues := make([]int, len(values))
	for i, v := range values {
		multipliedValues[i] = v * multiplier
	}
	return multipliedValues
}

func add(values []int, additive int) []int {
	addedValues := make([]int, len(values))
	for i, v := range values {
		addedValues[i] = v + additive
	}
	return addedValues
}
