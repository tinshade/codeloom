package routes

import (
	"fmt"
	"net/http"

	"github.com/tinshade/codeloom/internal/auth"
)

func CustomServerRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/login", auth.VerifyJWT(auth.HandleReqRes))
	mux.HandleFunc("/api/data", apiDataHandler)
	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the homepage!")
}

func apiDataHandler(w http.ResponseWriter, r *http.Request) {
	data := "Some data from the API"
	auth.HandleReqRes(w, r)

	fmt.Fprintln(w, data)
}
