package models

import "github.com/google/uuid"

type Person struct {
	Id   uuid.UUID `json:"Id" bson:"Id"`
	Name string    `json:"Name" bson:"Name"`
}

type CreatePersonRequest struct {
	Name string `json:"name"`
}

type CreatePersonResponse struct {
	NewPersonId uuid.UUID `json:"personId"`
}
