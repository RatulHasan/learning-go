package movieController

import (
	"encoding/json"
	"fmt"
	dbHelper "github.com/RatulHasan/learning-go/helpers"
	movieModel "github.com/RatulHasan/learning-go/models"
	"github.com/gorilla/mux"
	"net/http"
)

// GetAllMovies get all movies
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	// if Method is POST then create movie
	if r.Method == "POST" {
		CreateMovie(w, r)
		return
	}
	movies := dbHelper.GetMovieList()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// convert json to byte array
	movieByte, err := json.Marshal(movies)
	write, err := w.Write(movieByte)
	if err != nil {
		return
	}
	fmt.Println(write)
}

// GetMovie get movie by id
func GetMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieId := vars["movieId"]
	movie := dbHelper.GetMovie(movieId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// convert json to byte array
	movieByte, err := json.Marshal(movie)
	write, err := w.Write(movieByte)
	if err != nil {
		return
	}
	fmt.Println(write)
}

// CreateMovie create movie
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie movieModel.Netflix
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, movie := range movieModel.Movies {
		dbHelper.InsertMovie(movie)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// convert json to byte array
	movieByte, err := json.Marshal(movie)
	write, err := w.Write(movieByte)
	if err != nil {
		return
	}
	fmt.Println(write)
}

// MarkAsWatched update movie
func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieId := vars["movieId"]
	movie := dbHelper.MarkAsWatched(movieId)
	movieByte, _ := json.Marshal(movie)
	write, _ := w.Write(movieByte)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Println(write)
}

// DeleteMovie delete movie
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieId := vars["movieId"]
	dbHelper.DeleteMovie(movieId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Printf("Movie %s deleted\n", movieId)
}

// DeleteAllMovies delete all movies
func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	dbHelper.DeleteAllMovies()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Println("All movies deleted")
}
