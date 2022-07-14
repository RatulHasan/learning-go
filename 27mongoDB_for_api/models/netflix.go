package movieModel

import "go.mongodb.org/mongo-driver/bson/primitive"

type Netflix struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title,omitempty"`
	Director  string             `bson:"director,omitempty"`
	Actors    []Actor            `bson:"actors,omitempty"`
	Genre     []Genre            `bson:"genre,omitempty"`
	Rating    string             `bson:"rating,omitempty"`
	Watched   bool               `bson:"watched,omitempty"`
	CreatedAt string             `bson:"created_at,omitempty"`
	UpdatedAt string             `bson:"updated_at,omitempty"`
}

type Actor struct {
	Id   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name,omitempty"`
}

type Genre struct {
	Id   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name,omitempty"`
}

// Movies Fake data
var Movies = []Netflix{
	{
		Id:       primitive.NewObjectID(),
		Title:    "The Shawshank Redemption",
		Director: "Frank Darabont",
		Actors: []Actor{
			{Id: primitive.NewObjectID(), Name: "Tim Robbins"},
			{Id: primitive.NewObjectID(), Name: "Morgan Freeman"},
		},
		Genre: []Genre{
			{Id: primitive.NewObjectID(), Name: "Drama"},
		},
		Rating:    "9.3",
		Watched:   false,
		CreatedAt: "2020-01-01",
		UpdatedAt: "2020-01-01",
	},
	{
		Id:       primitive.NewObjectID(),
		Title:    "The Godfather",
		Director: "Francis Ford Coppola",
		Actors: []Actor{
			{Id: primitive.NewObjectID(), Name: "Marlon Brando"},
			{Id: primitive.NewObjectID(), Name: "Al Pacino"},
		},
		Genre: []Genre{
			{Id: primitive.NewObjectID(), Name: "Drama"},
		},
		Rating:    "9.2",
		Watched:   false,
		CreatedAt: "2020-01-01",
		UpdatedAt: "2020-01-01",
	},
}
