package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/tinshade/codeloom/internal/auth"
	"github.com/tinshade/codeloom/internal/helpers"
	"github.com/tinshade/codeloom/internal/routes"
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

func main() {
	var _, err = helpers.RegisterEnvVars(".env")
	if err != nil {
		fmt.Println("Something went wrong while starting up the server!")
	}
	auth.CheckForEnvVars()

	var isFirstRun string = os.Getenv("SHOULD_CREATE_COLLECTION")
	if isFirstRun == "true" {
		helpers.HandleInitialDBSetup()
	}

	CustomServer()
}
