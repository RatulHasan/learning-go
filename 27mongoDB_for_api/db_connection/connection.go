package db_connection

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const ConnectionString = "mongodb+srv://ratul:30XxuhGXECnXJxzE@learning-go.curri.mongodb.net/?retryWrites=true&w=majority"
const DBName = "learning-Go"
const CollectionName = "netflix"

// Collection Most Important:
var Collection *mongo.Collection

// Init Create a new MongoDB connection
func Init() {
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
