package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://dokan-app.test/wp-json/dokan/v1/stores?badge_id=5&page=1"

func main() {
	fmt.Println("The URL is:", myUrl)

	// Parse the URL
	result, err := url.Parse(myUrl)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Scheme:", result.Scheme)
	fmt.Println("Host:", result.Host)
	fmt.Println("Port:", result.Port())
	fmt.Println("Path:", result.Path)
	fmt.Println("RawQuery:", result.RawQuery)

	// Parse the query
	//query, err := url.ParseQuery(result.RawQuery)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return
	//}

	// or Parse the query
	query := result.Query()

	// Print the query
	fmt.Println("Query:", query)

	// Print the query values using the for loop
	for key, value := range query {
		fmt.Println("Key:", key)
		fmt.Println("Value:", value)
	}

	// Construct the URL
	// Remember to use the pass by reference for the url.URL
	newUrl := &url.URL{
		Scheme:   result.Scheme,
		Host:     result.Host,
		Path:     result.Path,
		RawQuery: result.RawQuery,
	}

	fmt.Println("New URL:", newUrl.String())
}
