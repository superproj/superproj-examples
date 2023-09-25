package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AuditMessage struct {
	Matcher   string        `json:"matcher"`
	Request   []interface{} `json:"request"`
	Result    bool          `json:"result"`
	Explains  [][]string    `json:"explains"`
	Timestamp int64         `json:"timestamp"`
}

func main() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://zero:zero(#)666@127.0.0.1:27017/zero?authSource=zero")

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Ping the MongoDB server to check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Get a handle to the "zero" database and "audit" collection
	collection := client.Database("zero").Collection("audit")

	// Insert a document
	message := AuditMessage{
		Matcher: "111",
		//Request   []interface{} `json:"request"`
		//Result    bool          `json:"result"`
		////Explains  [][]string    `json:"explains"`
		//Timestamp int64         `json:"timestamp"`
	}
	insertResult, err := collection.InsertOne(context.Background(), message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted document ID:", insertResult.InsertedID)

	// Find a document
	var result AuditMessage
	err = collection.FindOne(context.Background(), bson.M{"matcher": "111"}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Found document:", result)

	// Update a document
	update := bson.M{"$set": bson.M{"matcher": "222"}}
	updateResult, err := collection.UpdateOne(context.Background(), bson.M{"matcher": "111"}, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated", updateResult.ModifiedCount, "document(s)")

	// Delete a document
	deleteResult, err := collection.DeleteOne(context.Background(), bson.M{"matcher": "222"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted", deleteResult.DeletedCount, "document(s)")
}
