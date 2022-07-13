package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

const getUrl = "http://localhost:2000/get"

func main() {
	PerformGetRequest(getUrl)
}

func PerformGetRequest(url string) {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)

	fmt.Println("Status:", response.Status)
	fmt.Println("Headers:", response.Header)
	fmt.Println("Content type:", response.Header.Get("content-type"))
	fmt.Println("Content length:", response.ContentLength)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	//fmt.Println("Body:", string(body))
	var responseString strings.Builder
	responseString.Write(body)
	fmt.Println("Response:", responseString.String())
}
