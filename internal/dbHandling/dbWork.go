package dbHandling

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Likely going to use mongoDB

// DbWork is a function that allows a user to work with the database
func DbWork() error {
	fmt.Println("DbWork")
	return nil
}

// DBInit is a function that allows a user to initialize the database
func DBInit() error {
	fmt.Println("DBInit")
	// Initialize the database options
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://wfcornelissen:iMtEnuxpm6MAMCBI@prim.weghf.mongodb.net/?retryWrites=true&w=majority&appName=Prim").SetServerAPIOptions(serverAPI)
	// Set up the database connection
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	return nil
}
