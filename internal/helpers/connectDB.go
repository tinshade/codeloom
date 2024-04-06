package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/tinshade/codeloom/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var _, err = RegisterEnvVars(".env")

func CheckConnection(client *mongo.Client) bool {
	err := client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Println("Connection established successfully!")
	return true

}

func LoadFixtures(fixturePath string, collection *mongo.Collection) bool {
	data, err := ReadStructFromJSON(fixturePath)
	if err != nil {
		log.Fatal("Error reading from file")
		return false
	}

	var listOfItems models.Users
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
	}
	if err := json.Unmarshal(jsonData, &listOfItems); err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
	}
	users := listOfItems.User
	usersModel := []interface{}{}
	if len(users) > 0 {

		for _, person := range users {
			usersModel = append(usersModel, person)
		}
	}

	res, err := collection.InsertMany(context.Background(), usersModel)
	if err != nil {
		return false
	}
	print(res)
	return true
}

func CreateCollection(db_name string, collection_name string, client *mongo.Client) *mongo.Collection {
	//* Create collection
	var collection *mongo.Collection = client.Database(db_name).Collection(collection_name)
	return collection

}

func FindOne(query bson.D, collection *mongo.Collection) models.Any {

	// Find one document
	var result interface{}
	err := collection.FindOne(context.Background(), query).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No matching documents found")
			return nil
		}
		log.Fatal(err)
	}
	fmt.Println(result, "Check this")
	return result
}

func FindMany(query bson.D, collection *mongo.Collection) models.Any {
	//! EXAMPLE FILTER
	//* filter := bson.D{{"age", bson.D{{"$gte", 25}}}} // Example: Find people with age greater than or equal to 25

	// Find documents
	cursor, err := collection.Find(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	// Iterate over the cursor and decode each document

	//TODO: Change the model to be dynamic
	var results []interface{}
	for cursor.Next(context.Background()) {
		var result interface{}
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}

	// Check if there was an error during the cursor iteration
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println("This is many results:\t", results)
	return results
}

func CheckIfExists(query bson.D, collection *mongo.Collection) int {
	// Find one document
	var result interface{}
	err := collection.FindOne(context.Background(), query).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("Record does not exist")
			return 0
		}
		log.Fatal(err)
		return -1
	}
	return 1
}

func UpdateOne(query bson.D, updateFields bson.D, collection *mongo.Collection) (models.Any, error) {
	// Define update criteria
	update := bson.D{
		{"$set", updateFields},
	}

	// Perform the update
	_, err := collection.UpdateOne(context.Background(), query, update)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	document := FindOne(query, collection)
	return document, nil

}

func GetCollectionInstance(collection_name string, client *mongo.Client) *mongo.Collection {
	// Access the database and collection
	collection := client.Database(os.Getenv("DB_NAME")).Collection(collection_name)
	return collection
}

func UpdateMany(query bson.D, updateQuery bson.D, collection *mongo.Collection) (models.Any, error) {

	// Define update criteria
	update := bson.D{
		{"$set", updateQuery},
	}

	// Perform the update
	updateResult, err := collection.UpdateMany(context.Background(), query, update)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Print the number of documents updated
	fmt.Printf("Updated %d documents\n", updateResult.ModifiedCount)
	document := FindOne(query, collection)
	return document, nil
}

func DeleteRecords(deletionType string, query string, collection *mongo.Collection) bool {
	if deletionType == "many" {
		deleteResult, err := collection.DeleteMany(context.Background(), query)
		if err != nil {
			log.Fatal(err)
			return false
		}

		// Print the number of documents deleted
		fmt.Printf("Deleted %v document(s)\n", deleteResult.DeletedCount)
	} else {
		deleteResult, err := collection.DeleteOne(context.Background(), query)
		if err != nil {
			log.Fatal(err)
			return false
		}

		// Print the number of documents deleted
		fmt.Printf("Deleted %v document(s)\n", deleteResult.DeletedCount)
	}

	return true
}

var ClientInstance *mongo.Client = Connect()

func Connect() *mongo.Client {
	//* Find the DB string from .env file
	var connectionString string = os.Getenv("DB_CONNECTION_STRING")

	//* Connecting to the database
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
		Disconnect(client)
	}

	CheckConnection(client)
	return client
}

func Disconnect(client *mongo.Client) {
	defer client.Disconnect(context.Background())

}
