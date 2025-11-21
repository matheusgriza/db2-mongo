package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id          uuid.UUID
	Title       string
	Date        time.Time
	Invited     []uuid.UUID
	Description string
}

type CreateTaskRequest struct {
	Date        time.Time   `bson:"date" json:"date"`
	Invited     []uuid.UUID `bson:"invited" json:"invited`
	Title       string      `bson:"title" json:"title`
	Description string      `bson:"description" json:"description"`
}

type CreateTaskResponse struct {
	Id uuid.UUID `bson:"_id" json:"id"`
}

type UpdateTaskRequest struct {
	Title       string `bson:"title" json:"title`
	Description string `json:"description"`
}

type ManageInvitedTask struct {
	Ids []uuid.UUID `bson: "ids" json:"ids"`
}
