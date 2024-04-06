package models

//https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#Collection.InsertMany
import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Any interface{}

type ModelStruct struct {
	Data map[string]interface{}
}

type Users struct {
	User []User `json:"data"`
}

type User struct {
	Id             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	First_name     string             `bson:"first_name" json:"first_name"`
	Last_name      string             `bson:"last_name" json:"last_name"`
	Title          string             `bson:"title,omitempty" json:"title,omitempty"`
	Email          string             `bson:"email" json:"email"`
	Is_admin       bool               `bson:"is_admin" json:"is_admin"`
	Is_reviewer    bool               `bson:"is_reviewer" json:"is_reviewer"`
	Is_active      bool               `bson:"is_active" json:"is_active"`
	Can_merge      bool               `bson:"can_merge" json:"can_merge"`
	Token          string             `bson:"token" json:"token"`
	Token_end_date time.Time          `bson:"token_end_date" json:"token_end_date"`
	Password       string             `bson:"password" json:"password"`
	Group          Group              `bson:"group" json:"group"`
	Joined_on      time.Time          `bson:"joined_on" json:"joined_on"`
}

type Group struct {
	Id                primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Group_name        string             `bson:"group_name" json:"group_name"`
	Group_description string             `bson:"group_description" json:"group_description"`
	Group_permissions Permissions        `bson:"group_permissions" json:"group_permissions"`
}

type Permissions struct {
	Id                     primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Permission_name        string             `bson:"permission_name" json:"permission_name"`
	Permission_description string             `bson:"permission_description" json:"permission_description"`
	Is_active              bool               `bson:"is_active" json:"is_active"`
}

type TicketBuckets struct {
	Id                 primitive.ObjectID `bson:"_id,omitempty" json "_id,omitempty"`
	Bucket_name        string             `bson:"bucket_name" json "bucket_name"`
	Bucket_description string             `bson:"bucket_description,omitempty" json "bucket_description,omitempty"`
	Start_date         time.Time          `bson:"start_date" json "start_date"`
	End_date           time.Time          `bson:"end_date" json "end_date"`
	Tickets            Tickets            `bson:"tickets" json "tickets"`
}

type Tickets struct {
	Id                           primitive.ObjectID `bson:"_id,omitempty" json "_id,omitempty"`
	Ticket_id                    string             `bson:"ticket_id" json "ticket_id"`
	Ticket_name                  string             `bson:"ticket_name" json "ticket_name"`
	Ticket_description           string             `bson:"ticket_description" json "ticket_description"`
	Attachments                  string             `bson:"attachments,omitempty" json "attachments,omitempty"`
	Started_by                   User               `bson:"started_by" json "started_by"`
	Start_date                   time.Time          `bson:"start_date" json "start_date"`
	End_date                     time.Time          `bson:"end_date" json "end_date"`
	Assigned_to                  User               `bson:"assigned_to" json "assigned_to"`
	Notify_on_activity           bool               `bson:"notify_on_activity" json "notify_on_activity"`
	Notify_before_expiry         bool               `bson:"notify_before_expiry" json "notify_before_expiry"`
	Notify_on_expiry             bool               `bson:"notify_on_expiry" json "notify_on_expiry"`
	Sent_notification_for_expiry bool               `bson:"sent_notification_for_expiry" json "sent_notification_for_expiry"`
}
