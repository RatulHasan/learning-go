package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/posts"

func main() {
	// Create a new request using GET method
	response, err := http.Get(url)
	if err != nil {
		return
	}

	//defer response.Body.Close()

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	// Print the response type
	fmt.Printf("The response type is %T \n", response)

	// Print the status
	fmt.Printf("Status: %s\n", response.Status)

	// Print the headers
	fmt.Printf("Headers: %v\n", response.Header)

	// Print the body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	// Print raw body
	fmt.Println("Body:", body)

	// Print the body as string
	fmt.Println("Body:", string(body))
}
