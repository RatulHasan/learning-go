package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("Post form data")
	data := url.Values{}
	data.Add("name", "John")
	data.Add("age", "30")
	data.Add("email", "jhon@doe.com")
	data.Add("password", "123456")
	data.Add("password_confirmation", "123456")
	data.Add("location", "New York")
	PerformPostForm("http://localhost:2000/postform", data)
}

func PerformPostForm(url string, data url.Values) {
	response, err := http.PostForm(url, data)
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
	fmt.Println("Content type:", response.Header.Get("content-type"))
	fmt.Println("Content length:", response.ContentLength)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Body:", string(body))
}
