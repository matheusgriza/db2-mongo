package repositories

import (
	"task-api/internal/models"
	"task-api/internal/repositories/persons"

	"github.com/google/uuid"
)

type Repositories struct {
	Person interface {
		GetPerson(id uuid.UUID) *models.Person
		GetAllPerson() []models.Person
		AddPerson(newPerson models.Person)
	}
}

func New() *Repositories {
	return &Repositories{
		Person: persons.New(),
	}
}
