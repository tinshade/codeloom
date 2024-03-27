package main

import (
	"fmt"
	"net/http"

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

var data, err = helpers.RegisterEnvVars(".env.example")

func main() {
	auth.CheckForEnvVars()
	fmt.Println(data, err)
	if err != nil {
		CustomServer()
	}
}
