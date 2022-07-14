package db_connection

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

// Collection Most Important:
var Collection *mongo.Collection

// Init Create a new MongoDB connection
func Init() {
	// Load local env variables
	err := godotenv.Load("./local.env")
	if err != nil {
		log.Fatal("Error loading local.env file")
	}
	ConnectionString := os.Getenv("ConnectionString")
	DBName := os.Getenv("DBName")
	CollectionName := os.Getenv("CollectionName")

	// Client options
	clientOptions := options.Client().ApplyURI(ConnectionString)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	// Check the connection
	err = client.Ping(nil, nil)
	if err != nil {
		panic(err)
	}

	// Get a handle for your collection
	Collection = client.Database(DBName).Collection(CollectionName)

	fmt.Println("Connected to MongoDB!")
}
