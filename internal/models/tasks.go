package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id          uuid.UUID
	date        time.Time
	description string
	invited     []Person
}

type CreateTaskRequest struct {
	date        time.Time `json:date`
	description string    `json:description`
	invited     []Person  `json:invited`
}

type CreateTaskResponse struct {
	Id uuid.UUID
}
