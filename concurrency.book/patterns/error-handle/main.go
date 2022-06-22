package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Bad Example: ")
	badExample()

	fmt.Println("Good Exmple: ")
	goodExample()
}

func badExample() {
	done := make(chan interface{})
	defer close(done)
	urls := []string{"https://www.google.com", "https://badhost"}
	for response := range checkStatusBad(done, urls...) {
		fmt.Printf("Response: %v\n", response.Status)
	}
}

func checkStatusBad(done <-chan interface{}, urls ...string) <-chan *http.Response {
	responses := make(chan *http.Response)
	go func() {
		defer close(responses)
		for _, url := range urls {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
				continue
			}

			select {
			case <-done:
				return
			case responses <- resp:
			}
		}
	}()
	return responses
}

type Result struct {
	Error    error
	Response *http.Response
}

func goodExample() {
	done := make(chan interface{})
	defer close(done)
	urls := []string{"https://www.google.com", "https://badhost"}
	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			fmt.Printf("error: %v\n", result.Error)
			continue
		}
		fmt.Printf("Response: %v\n", result.Response.Status)
	}
}

func checkStatus(done <-chan interface{}, urls ...string) <-chan Result {
	responses := make(chan Result)
	go func() {
		defer close(responses)
		for _, url := range urls {
			resp, err := http.Get(url)
			result := Result{Error: err, Response: resp}
			select {
			case <-done:
				return
			case responses <- result:
			}
		}
	}()
	return responses
}
