package dbHelper

import (
	"context"
	"fmt"
	"github.com/RatulHasan/learning-go/db_connection"
	movieModel "github.com/RatulHasan/learning-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

// GetMovieList get list of movies
func GetMovieList() []movieModel.Netflix {
	var movies []movieModel.Netflix
	cur, err := db_connection.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.Background()) {
		var movie movieModel.Netflix
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(cur, context.Background())

	return movies
}

// GetMovie get one movie from the collection
func GetMovie(movieId string) movieModel.Netflix {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	var movie movieModel.Netflix
	err := db_connection.Collection.FindOne(context.Background(), filter).Decode(&movie)
	if err != nil {
		log.Fatal(err)
	}

	return movie
}

// InsertMovie insert one movie into the collection
func InsertMovie(movie movieModel.Netflix) {
	// Insert one movie.
	result, err := db_connection.Collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", result.InsertedID)
}

// MarkAsWatched if the movie is watched.
func MarkAsWatched(movieId string) movieModel.Netflix {
	id, _ := primitive.ObjectIDFromHex(movieId)
	updateId := bson.M{"_id": id}
	data := bson.M{"$set": bson.M{"watched": true}}
	result, err := db_connection.Collection.UpdateOne(context.Background(), updateId, data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", result.MatchedCount, result.ModifiedCount)
	return GetMovie(movieId)
}

// DeleteMovie delete one movie from the collection
func DeleteMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	deleteId := bson.M{"_id": id}
	result, err := db_connection.Collection.DeleteOne(context.Background(), deleteId)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted %v documents.\n", result.DeletedCount)
}

// DeleteAllMovies delete all movies from the collection
func DeleteAllMovies() {
	result, err := db_connection.Collection.DeleteMany(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted %v documents.\n", result.DeletedCount)
}
