package usecases

import (
	"context"
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

func (u UseCases) GetPerson(ctx context.Context, id uuid.UUID) (*models.Person, error) {
	person, err := u.repos.Person.GetPerson(ctx, id)
	if err != nil {
		return nil, err
	}

	return person, nil
}

func (u UseCases) GetAllPerson(ctx context.Context) ([]models.Person, error) {
	persons, err := u.repos.Person.GetAllPerson(ctx)
	if err != nil {
		return nil, err
	}
	return persons, nil
}

func (u UseCases) AddPerson(ctx context.Context, newPerson models.CreatePersonRequest) (uuid.UUID, error) {
	repoReq := models.Person{
		Id:   uuid.New(),
		Name: newPerson.Name,
	}

	err := u.repos.Person.AddPerson(ctx, repoReq)
	if err != nil {
		return uuid.Nil, err
	}
	return repoReq.Id, nil
}
