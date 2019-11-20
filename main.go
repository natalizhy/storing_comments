package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go.mongodb.org/mongo-driver/bson"
	"github.com/go.mongodb.org/mongo-driver/mongo"
	"github.com/go.mongodb.org/mongo-driver/mongo/options"
	"github.com/go.mongodb.org/mongo-driver/mongo/readpref"
 	"github.com/natalizhy/storing comments/assets/mongodb"
 	//"github.com/mongodb/mongo-go-driver"

)

// You will be using this Trainer type later in the program
type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	// Rest of the code will go here
	//CreateMongoClient()
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}

	// Create connect
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	collection := client.Database("test").Collection("trainers")
	fmt.Println(collection)

	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
