package helpers

import (
	"os"
	"fmt"
	"log"
	"context"
	"encoding/json"
	"github.com/tinshade/codeloom/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func CheckConnection(client *mongo.Client) bool{
	err := client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Println("Connection established successfully!")
	return true

}



func LoadFixtures(fixturePath string, collection *mongo.Collection) bool{
	data,err := ReadStructFromJSON(fixturePath)
	if err != nil{
		log.Fatal("Error reading from file")
		return false
	}
	
	var listOfItems models.Users
	jsonData, err := json.Marshal(data)
	if err != nil{
		fmt.Println("Error marshaling JSON:", err)
	}
	if err := json.Unmarshal(jsonData, &listOfItems); err != nil {
        fmt.Println("Error unmarshaling JSON:", err)
    }
	users := listOfItems.User
	usersModel := []interface{}{}
	
	for _, person := range users {
        usersModel = append(usersModel, person)
    }

	res, err := collection.InsertMany(context.Background(), usersModel)
	if err != nil { return false }
	print(res)
	return true
}


func CreateCollection(db_name string, collection_name string, client *mongo.Client ) *mongo.Collection{
	//* Create collection
	var collection *mongo.Collection = client.Database(db_name).Collection(collection_name)
	return collection

}

func Connect() *mongo.Client{
	//* Find the DB string from .env file
	var connectionString string = os.Getenv("DB_CONNECTION_STRING")
	fmt.Println(connectionString)

	//* Connecting to the database
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil{
		log.Fatal(err)
	}

	CheckConnection(client)
	
	
	return client
}