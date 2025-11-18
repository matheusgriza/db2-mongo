package usecases

import (
	"task-api/internal/models"
	"task-api/internal/repositories"

	"github.com/google/uuid"
)

type UseCases struct {
	repos *repositories.Repositories
}

func New(repos *repositories.Repositories) *UseCases {
	return &UseCases{
		repos: repos,
	}
}

func (u UseCases) GetPerson(id uuid.UUID) *models.Person {
	person := u.repos.Person.GetPerson(id)
	if person == nil {
		panic("User not found")
	}
	return person
}

func (u UseCases) GetAllPerson() []models.Person {
	persons := u.repos.Person.GetAllPerson()
	return persons
}

func (u UseCases) AddPerson(newPerson models.CreatePersonRequest) (uuid.UUID, error) {
	repoReq := models.Person{
		Id:   uuid.New(),
		Name: newPerson.Name,
	}
	u.repos.Person.AddPerson(repoReq)

	return repoReq.Id, nil
}
