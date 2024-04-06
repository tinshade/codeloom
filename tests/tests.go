package tests

import (
	"fmt"
	// "github.com.tinshade/codeloom/cmd/codeloom/main"
	"github.com/tinshade/codeloom/internal/helpers"
	"github.com/tinshade/codeloom/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func findMany(query bson.D, collection *mongo.Collection) {
	result := helpers.FindMany(query, collection)

	//USING THE MANY RESULTS THAT YOU GET
	if data, ok := result.([]models.User); ok {
		for _, item := range data {
			fmt.Printf("Name: %s, last Name: %s, Email: %s\n", item.First_name, item.Last_name, item.Email)
		}
	} else {
		fmt.Println("err")
	}
}

func findOne(query bson.D, collection *mongo.Collection) {
	//Finding One Result
	one_res := helpers.FindOne(query, collection)
	fmt.Println("one res: " + fmt.Sprintf("%v", one_res))
}

func deleteRecords(query bson.D, collection *mongo.Collection) {
	helpers.DeleteRecords(query, collection)
}

func DBTests() {
	var CollectionInstance *mongo.Collection = helpers.GetCollectionInstance("users", main.ClientInstance)
	query := bson.D{{"is_admin", true}}
	deletionQuery := bson.D{{"first_name", "Lari2"}}
	findMany(query, CollectionInstance)
	findOne(query, CollectionInstance)
	deleteRecords(deletionQuery, CollectionInstance)
}
