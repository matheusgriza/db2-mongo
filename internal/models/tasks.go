package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id          uuid.UUID   `bson:"Id" json:"id"`
	Title       string      `bson:"Title" json:"title"`
	Date        time.Time   `bson:"Date" json:"date"`
	Invited     []uuid.UUID `bson:"Invited" json:"invited"`
	Description string      `bson:"Description" json:"description"`
}

type CreateTaskRequest struct {
	Date        time.Time   `bson:"Date" json:"date"`
	Invited     []uuid.UUID `bson:"Invited" json:"invited"`
	Title       string      `bson:"Title" json:"title"`
	Description string      `bson:"Description" json:"description"`
}

type CreateTaskResponse struct {
	Id uuid.UUID `bson:"Id" json:"id"`
}

type UpdateTaskRequest struct {
	Title       string `bson:"Title" json:"title"`
	Description string `bson:"Description" json:"description"`
}

type ManageTaskInvited struct {
	Invited []uuid.UUID `json:"ids"`
}
