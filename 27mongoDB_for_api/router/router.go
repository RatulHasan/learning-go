package router

import (
	movieController "github.com/RatulHasan/learning-go/controllers"
	"github.com/RatulHasan/learning-go/db_connection"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Routes() {
	// Connect to MongoDB
	db_connection.Init()

	r := mux.NewRouter()
	r.HandleFunc("/", movieController.GetAllMovies).Methods("GET", "POST")
	//r.HandleFunc("/", movieController.CreateMovie).Methods("POST")
	r.HandleFunc("/{movieId}", movieController.GetMovie).Methods("GET")
	r.HandleFunc("/{movieId}", movieController.MarkAsWatched).Methods("PUT")
	r.HandleFunc("/{movieId}", movieController.DeleteMovie).Methods("DELETE")
	r.HandleFunc("/", movieController.DeleteAllMovies).Methods("DELETE")

	// Listen and serve on port 2000
	log.Fatal(http.ListenAndServe(":2000", r))
}
