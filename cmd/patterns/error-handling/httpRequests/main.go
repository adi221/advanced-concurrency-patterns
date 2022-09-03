package main

import (
	"fmt"
	"net/http"
)

type Result struct {
	Error    error
	Response *http.Response
}

func checkStatus(done <-chan interface{}, urls ...string) <-chan Result {
	results := make(chan Result)
	go func() {
		defer close(results)

		for _, url := range urls {
			resp, err := http.Get(url)
			r := Result{Error: err, Response: resp}
			select {
			case <-done:
				return
			case results <- r:
			}
		}
	}()
	return results
}

func main() {
	done := make(chan interface{})
	defer close(done)
	var errCount int
	urls := []string{"https://www.google.com", "https://badhost", "a", "b", "c"}

	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			fmt.Printf("error %v\n", result.Error)
			errCount++
			if errCount > 2 {
				fmt.Println("Too many errors, quitting")
				break
			}
			continue
		}
		fmt.Printf("Result: %v\n", result.Response)
	}
}

// func main() {
// 	done := make(chan interface{})
// 	defer close(done)

// 	urls := []string{"https://www.google.com", "https://badhost", "https://www.facebook.com"}
// 	for res := range checkStatus(done, urls...) {
// 		if res.Error != nil {
// 			fmt.Printf("error %v\n", res.Error)
// 			continue
// 		}
// 		fmt.Printf("Response %v\n", res.Response)
// 	}
// }
