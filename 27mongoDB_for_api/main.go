package main

import (
	"github.com/RatulHasan/learning-go/db_connection"
	"github.com/RatulHasan/learning-go/router"
)

func init() {
	router.Routes()
}

func main() {
	// Connect to MongoDB
	db_connection.Init()
}
