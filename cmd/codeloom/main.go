package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/tinshade/codeloom/internal/auth"
	"github.com/tinshade/codeloom/internal/helpers"
	"github.com/tinshade/codeloom/internal/routes"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CustomServer() {
	router := routes.CustomServerRoutes()
	port := 8080
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf(("Starting server on %s"), addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}

var _, err = helpers.RegisterEnvVars(".env")

func handleInitialDBSetup() {
	var DBName string = os.Getenv("DB_NAME")
	var CollectionName string = os.Getenv("COLLECTION_NAME")
	var collection *mongo.Collection = helpers.CreateCollection(DBName, CollectionName, helpers.ClientInstance)

	//* LOADING FIXTURES FROM PATH
	var usersFixturePath string = fmt.Sprintf("%s/users.json", helpers.FIXTURES_BASE_PATH)

	helpers.LoadFixtures(usersFixturePath, collection)
}

func main() {
	auth.CheckForEnvVars()
	if err != nil {
		fmt.Println("Something went wrong while starting up the server!")
	}

	var isFirstRun string = os.Getenv("SHOULD_CREATE_COLLECTION")
	if isFirstRun == "true" {
		handleInitialDBSetup()
	}

	CustomServer()
}
