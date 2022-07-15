package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}
var waitGroup sync.WaitGroup
var mutex sync.Mutex

func main() {
	websiteList := []string{
		"http://localhost:2000/",
		"https://jsonplaceholder.typicode.com/posts",
		"https://jsonplaceholder.typicode.com/todos",
		"https://jsonplaceholder.typicode.com/users",
		"https://jsonplaceholder.typicode.com/comments",
		"https://jsonplaceholder.typicode.com/albums",
		"https://fb.com",
		"https://google.com",
		"https://youtube.com",
		"https://twitter.com",
		"https://instagram.com",
		"https://linkedin.com",
		"https://pinterest.com",
		"https://tumblr.com",
		"https://github.com",
		"https://stackoverflow.com",
	}

	// Example of Concurrency.
	//for _, website := range websiteList {
	//	GetStatusCodeConcurrency(website)
	//}

	// Example of Goroutines or Parallelism.
	// GoRoutines are more efficient and faster than Concurrency.
	for _, website := range websiteList {
		go GetStatusCodeGoroutines(website)
		waitGroup.Add(1)
	}
	waitGroup.Wait()

	fmt.Println("All websites are done!")
}

// GetStatusCodeConcurrency is a function that returns the status code of a response. Example of Concurrency.
func GetStatusCodeConcurrency(endpoint string) {
	response, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}

	fmt.Println("Status:", response.Status, "for", endpoint)
}

// GetStatusCodeGoroutines is a function that returns the status code of a response. Example of Goroutines or Parallelism.
func GetStatusCodeGoroutines(endpoint string) {
	defer waitGroup.Done()
	response, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}

	mutex.Lock()
	signals = append(signals, endpoint)
	mutex.Unlock()
	fmt.Println("Status:", response.Status, "for", endpoint)
}
