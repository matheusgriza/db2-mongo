package models

import "github.com/google/uuid"

type Person struct {
	Id   uuid.UUID
	Name string
}

type CreatePersonRequest struct {
	Name string
}
