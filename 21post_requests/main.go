package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

const postUrl = "http://localhost:2000/post"

func main() {
	PerformPostJsonRequest(postUrl)
}

// PerformPostJsonRequest Make post request.
func PerformPostJsonRequest(url string) {
	// Fake json payload data
	formJsonData := strings.NewReader(`{"name":"John Doe","age":30, "city":"New York", "country":"USA", "email":"john@doe.com", "phone":"123456789"}`)

	// Make request
	response, err := http.Post(url, "application/json", formJsonData)
	if err != nil {
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)

	fmt.Println("Status:", response.Status)

	fmt.Println("Headers:", response.Header)
	bytesContent, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	fmt.Println("Body:", string(bytesContent))
}
