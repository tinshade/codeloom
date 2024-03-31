package models 
//https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#Collection.InsertMany
import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Any interface{}

type Users struct {
	User []User `json:"data"`
}

type User struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	First_name string `bson:"first_name"`
	Last_name string `bson:"last_name"`
	Title string `bson:"title,omitempty"`
	Email string `bson:"email"`
	Is_admin bool `bson:"is_admin"`
	Is_reviewer bool `bson:"is_reviewer"`
	Is_active bool `bson:"is_active"`
	Can_merge bool `bson:"can_merge"`
	Token string `bson:"token"`
	Token_end_date time.Time `bson:"token_end_date"`
	Password string `bson:"password"`
	Group Group `bson:"group"`
	Joined_on time.Time
}


type Group struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Group_name string `bson:"group_name"`
	Group_description string `bson:"group_description"`
	Group_permissions Permissions `bson:"group_permissions"`
}


type Permissions struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Permission_name string `bson:"permission_name"`
	Permission_description string `bson:"permission_description"`
	Is_active bool `bson:"is_active"`
}


type TicketBuckets struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Bucket_name string `bson:"bucket_name"`
	Bucket_description string `bson:"bucket_description,omitempty"`
	Start_date time.Time `bson:"start_date"`
	End_date time.Time `bson:"end_date"`
	Tickets Tickets `bson:"tickets"`
}


type Tickets struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Ticket_id string `bson:"ticket_id"`
	Ticket_name string `bson:"ticket_name"`
	Ticket_description string `bson:"ticket_description"`
	Attachments string `bson:"attachments,omitempty"`
	Started_by User `bson:"started_by"`
	Start_date time.Time `bson:"start_date"`
	End_date time.Time `bson:"end_date"`
	Assigned_to User `bson:"assigned_to"`
	Notify_on_activity bool `bson:"notify_on_activity"`
	Notify_before_expiry bool `bson:"notify_before_expiry"`
	Notify_on_expiry bool `bson:"notify_on_expiry"`
	Sent_notification_for_expiry bool `bson:"sent_notification_for_expiry"`
}