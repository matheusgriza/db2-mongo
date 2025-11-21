package usecases

import (
	"context"
	"errors"
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

// should split it in different files

func (u UseCases) GetTask(ctx context.Context, id uuid.UUID) (*models.Task, error) {
	task, err := u.repos.Task.GetTask(ctx, id)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (u UseCases) GetAllTask(ctx context.Context) ([]models.Task, error) {
	task, err := u.repos.Task.GetAllTask(ctx)

	if err != nil {
		return nil, err
	}

	return task, nil

}

func (u UseCases) AddTask(ctx context.Context, newTask models.Task) (uuid.UUID, error) {
	taskReq := models.Task{
		Id:          uuid.New(),
		Description: newTask.Description,
		Date:        newTask.Date,
		Invited:     newTask.Invited,
	}

	valid, err := u.repos.Person.ValidateUUID(ctx, taskReq.Invited)

	if err != nil {
		return uuid.Nil, err
	}

	if !valid {
		return uuid.Nil, errors.New("One or more invited person IDs do not exist")
	}

	u.repos.Task.AddTask(ctx, taskReq)

	return taskReq.Id, nil
}

func (u UseCases) UpdateTask(ctx context.Context, task models.UpdateTaskRequest) (uuid.UUID, error) {
	// 4 - Alterar titulo e descriçao em um compromisso
	return uuid.Nil, nil

}

func (u UseCases) addInvited(ctx context.Context, id []uuid.UUID) (uuid.UUID, error) {
	return uuid.Nil, nil
}

func (u UseCases) removeInvited(ctx context.Context, id []uuid.UUID) (uuid.UUID, error) {
	return uuid.Nil, nil
}
