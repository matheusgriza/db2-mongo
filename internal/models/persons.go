package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Person struct {
	Id   primitive.ObjectID `json:"Id" bson:"_id"`
	Name string             `json:"Name" bson:"Name"`
}

type CreatePersonRequest struct {
	Name string `json:"name"`
}

type CreatePersonResponse struct {
	NewPersonId primitive.ObjectID `json:"personId"`
}

type GetTaskPerson struct {
	Id   primitive.ObjectID `json:"Id" bson:"_id"`
	Name string             `json:"Name" bson:"Name"`
}
