package apis

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tinshade/codeloom/internal/helpers"
	"github.com/tinshade/codeloom/internal/models"
	"go.mongodb.org/mongo-driver/mongo"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// REGISTER
var clientInstance *mongo.Client = helpers.ClientInstance
var UsersCollection *mongo.Collection = helpers.GetCollectionInstance("users", clientInstance)

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		return Register(w, r, UsersCollection, ``)
	}
	return nil
}

func Register(w http.ResponseWriter, r *http.Request, collection *mongo.Collection, data string) error {
	var user models.User
	err := json.Unmarshal([]byte(data), &user)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Insert the document
	result, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(result)

	return nil
}

// LOGIN
func Login() {}

// FORGOT PASSWORD
func ForgotPassword() {}

// RESET PASSWORD
func ResetPassword() {}

// AUTHORIZE FOR PR ACTIONS
func AuthorizePermissions() {}
