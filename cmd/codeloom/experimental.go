package main

import (
	"context"
	"fmt"
	"github.com/tinshade/codeloom/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func FindOne() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Access the database and collection
	collection := client.Database("codeloom").Collection("users")

	// Define filter criteria
	filter := bson.D{{"is_admin", true}}

	// Find one document
	var result models.User
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No matching documents found")
			return
		}
		log.Fatal(err)
	}

	// Print the result
	fmt.Printf("Found document:%+v\n", result)
}

func FindMany() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Access the database and collection
	collection := client.Database("codeloom").Collection("users")

	// Define filter criteria
	// filter := bson.D{{"age", bson.D{{"$gte", 25}}}} // Example: Find people with age greater than or equal to 25
	filter := bson.D{{"is_admin", true}}

	// Find documents
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	// Iterate over the cursor and decode each document
	var results []models.User
	for cursor.Next(context.Background()) {
		var result models.User
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}

	// Check if there was an error during the cursor iteration
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	// Print the results
	fmt.Printf("%T", results)
	fmt.Println("Found documents:")
	for _, result := range results {
		fmt.Printf("Name: %s, Age: %d, Email: %s\n", result.First_name, result.Last_name, result.Email)
	}

	fmt.Println()
}
func CheckIfExists() bool {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Access the database and collection
	collection := client.Database("codeloom").Collection("users")

	// Define filter criteria to check if the record exists
	filter := bson.D{{"first_name", "Larissa"}}

	// Find one document
	var result models.User
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("Record does not exist")
			return false
		}
		log.Fatal(err)
		return false
	}

	// Print the result
	fmt.Printf("\n\n\nRecord found: %+v\n\n\n\n", result)
	return true
}
func UpdateOne() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Access the database and collection
	collection := client.Database("codeloom").Collection("users")

	// Define filter criteria to find the document to update
	filter := bson.D{{"last_name", "Pinto"}}

	// Define update criteria
	update := bson.D{
		{"$set", bson.D{
			{"last_name", "Iyengar"},
			// {"email", "john@example.com"},
		}},
	}

	// Perform the update
	updateResult, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	// Print the number of documents updated
	fmt.Printf("Updated %d document\n", updateResult.ModifiedCount)

}

func DeleteRecords(deletionType string) bool {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Access the database and collection
	collection := client.Database("codeloom").Collection("users")

	// Define filter criteria to identify the document to delete
	filter := bson.D{{"first_name", "Lari2"}} // Example: Delete a document with name "John"

	// Perform the delete operation

	if deletionType == "many" {
		deleteResult, err := collection.DeleteMany(context.Background(), filter)
		if err != nil {
			log.Fatal(err)
			return false
		}

		// Print the number of documents deleted
		fmt.Printf("Deleted %v document(s)\n", deleteResult.DeletedCount)
	} else {
		deleteResult, err := collection.DeleteOne(context.Background(), filter)
		if err != nil {
			log.Fatal(err)
			return false
		}

		// Print the number of documents deleted
		fmt.Printf("Deleted %v document(s)\n", deleteResult.DeletedCount)
	}

	return true
}
