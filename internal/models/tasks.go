package models

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	Id          primitive.ObjectID `bson:"Id" json:"id"`
	Title       string             `bson:"Title" json:"title"`
	Date        time.Time          `bson:"Date" json:"date"`
	Invited     []uuid.UUID        `bson:"Invited" json:"invited"`
	Description string             `bson:"Description" json:"description"`
}

type CreateTaskRequest struct {
	Date        time.Time   `bson:"Date" json:"date"`
	Invited     []uuid.UUID `bson:"Invited" json:"invited"`
	Title       string      `bson:"Title" json:"title"`
	Description string      `bson:"Description" json:"description"`
}

type CreateTaskResponse struct {
	Id primitive.ObjectID `bson:"Id" json:"id"`
}

type UpdateTaskRequest struct {
	Title       string `bson:"Title" json:"title"`
	Description string `bson:"Description" json:"description"`
}

type ManageTaskInvited struct {
	Invited []primitive.ObjectID `json:"ids"`
}
