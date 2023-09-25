package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type TLSOptions struct {
	// UseTLS specifies whether should be encrypted with TLS if possible.
	UseTLS             bool   `json:"use-tls" mapstructure:"use-tls"`
	InsecureSkipVerify bool   `json:"insecure-skip-verify" mapstructure:"insecure-skip-verify"`
	CaCert             string `json:"ca-cert" mapstructure:"ca-cert"`
	Cert               string `json:"cert" mapstructure:"cert"`
	Key                string `json:"key" mapstructure:"key"`
}

type MongoOptions struct {
	// options needed when connnect to mongo db.
	TLSOptions *TLSOptions   `json:"tls" mapstructure:"tls"`
	URL        string        `json:"url" mapstructure:"url"`
	Timeout    time.Duration `json:"timeout" mapstructure:"timeout"`

	Database   string `json:"database" mapstructure:"database"`
	Collection string `json:"collection" mapstructure:"collection"`
	Username   string `json:"username" mapstructure:"username"`
	Password   string `json:"password" mapstructure:"password"`
}

type AuditMessage struct {
	Matcher   string        `json:"matcher"`
	Request   []interface{} `json:"request"`
	Result    bool          `json:"result"`
	Explains  [][]string    `json:"explains"`
	Timestamp int64         `json:"timestamp"`
}

func main() {
	opts := &MongoOptions{
		URL:        "mongodb://127.0.0.1:27017",
		Timeout:    3 * time.Second,
		Database:   "zero",
		Collection: "audit",
		Username:   "zero",
		Password:   "zero(#)666",
	}
	/*
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
	*/
	client, err := opts.NewClient()
	if err != nil {
		panic(err)
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

func (o *MongoOptions) NewClient() (*mongo.Client, error) {
	// Set client options
	opts := options.Client().
		SetConnectTimeout(o.Timeout).
		SetSocketTimeout(o.Timeout).
		SetServerSelectionTimeout(o.Timeout).
		ApplyURI(o.URL)

	if o.Username != "" && o.Password != "" {
		creds := options.Credential{
			AuthSource: o.Database,
			Username:   o.Username,
			Password:   o.Password,
		}
		opts.SetAuth(creds)
	}

	if o.TLSOptions != nil {
		//opts.TLSConfig = o.TLSOptions.TLSConfig()

	}

	if opts.ReadPreference == nil {
		opts.ReadPreference = readpref.Nearest()
	}

	ctx, cancel := context.WithTimeout(context.Background(), o.Timeout)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	// Ping the MongoDB server to check the connection
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return client, nil
}
