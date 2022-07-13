package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

const getUrl = "https://jsonplaceholder.typicode.com/posts"

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

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	// Decode the JSON data into a slice of structs.
	data := jsonDecode(body)

	// Print the data.
	for _, post := range data {
		fmt.Printf("UserId: %d \n", post.UserId)
		fmt.Printf("Id: %d \n", post.Id)
		fmt.Printf("Title: %s \n", post.Title)
		fmt.Printf("Body: %s \n", post.Body)
	}
}

func jsonDecode(jsonData []byte) []Post {
	var posts []Post
	err := json.Unmarshal(jsonData, &posts)
	if err != nil {
		panic(err)
	}

	return posts
}
