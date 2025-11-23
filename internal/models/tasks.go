package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	Id          primitive.ObjectID   `bson:"_id" json:"id"`
	Title       string               `bson:"Title" json:"title"`
	Date        time.Time            `bson:"Date" json:"date"`
	Invited     []primitive.ObjectID `bson:"Invited" json:"invited"`
	Description string               `bson:"Description" json:"description"`
}

type CreateTaskRequest struct {
	Date        time.Time            `bson:"Date" json:"date"`
	Invited     []primitive.ObjectID `bson:"Invited" json:"invited"`
	Title       string               `bson:"Title" json:"title"`
	Description string               `bson:"Description" json:"description"`
}

type CreateTaskResponse struct {
	Id primitive.ObjectID `bson:"_id" json:"id"`
}

type UpdateTaskRequest struct {
	Title       string `bson:"Title" json:"title"`
	Description string `bson:"Description" json:"description"`
}

type ManageTaskInvited struct {
	Invited []primitive.ObjectID `bson:"_id" json:"ids"`
}

type TaskGet struct {
	Id          primitive.ObjectID `bson:"_id" json:"id"`
	Title       string             `bson:"Title" json:"title"`
	Date        time.Time          `bson:"Date" json:"date"`
	Invited     []GetTaskPerson    `bson:"Invited" json:"invited"`
	Description string             `bson:"Description" json:"description"`
}
