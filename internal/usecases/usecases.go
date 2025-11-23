package usecases

import (
	"context"
	"errors"
	"task-api/internal/models"
	"task-api/internal/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UseCases struct {
	repos *repositories.Repositories
}

func New(repos *repositories.Repositories) *UseCases {
	return &UseCases{
		repos: repos,
	}
}

func (u UseCases) GetPerson(ctx context.Context, id primitive.ObjectID) (*models.Person, error) {
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

func (u UseCases) AddPerson(ctx context.Context, newPerson models.CreatePersonRequest) (primitive.ObjectID, error) {
	repoReq := models.Person{
		Id:   primitive.NewObjectID(),
		Name: newPerson.Name,
	}

	err := u.repos.Person.AddPerson(ctx, repoReq)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return repoReq.Id, nil
}

// should split it in different files
func (u UseCases) GetTask(ctx context.Context, id primitive.ObjectID) (*models.TaskGet, error) {
	task, err := u.repos.Task.GetTask(ctx, id)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (u UseCases) DeleteTask(ctx context.Context, id primitive.ObjectID) (*models.Task, error) {
	task, err := u.repos.Task.DeleteTask(ctx, id)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (u UseCases) GetAllTask(ctx context.Context) ([]models.TaskGet, error) {
	task, err := u.repos.Task.GetAllTask(ctx)

	if err != nil {
		return nil, err
	}

	return task, nil

}

func (u UseCases) AddTask(ctx context.Context, newTask models.CreateTaskRequest) (primitive.ObjectID, error) {
	taskReq := models.Task{
		Id:          primitive.NewObjectID(),
		Title:       newTask.Title,
		Description: newTask.Description,
		Date:        newTask.Date,
		Invited:     newTask.Invited,
	}

	valid, err := u.repos.Person.ValidateUUID(ctx, taskReq.Invited)

	if err != nil {
		return primitive.NilObjectID, err
	}

	if !valid {
		return primitive.NilObjectID, errors.New("one or more invited person IDs do not exist")
	}

	u.repos.Task.AddTask(ctx, taskReq)

	return taskReq.Id, nil
}

func (u UseCases) UpdateTask(ctx context.Context, id primitive.ObjectID, task models.UpdateTaskRequest) (primitive.ObjectID, error) {
	updateReq := models.UpdateTaskRequest{
		Title:       task.Title,
		Description: task.Description,
	}
	err := u.repos.Task.UpdateTask(ctx, id, updateReq)

	if err != nil {
		return primitive.NilObjectID, err
	}

	return id, nil

}

func (u UseCases) AddInvited(ctx context.Context, task primitive.ObjectID, ids []primitive.ObjectID) (primitive.ObjectID, error) {
	valid, err := u.repos.Person.ValidateUUID(ctx, ids)

	if err != nil {
		return primitive.NilObjectID, err
	}

	if !valid {
		return primitive.NilObjectID, errors.New("one or more invited person IDs do not exist or are already in the list")
	}

	u.repos.Task.AddInvited(ctx, task, ids)
	return task, nil
}

func (u UseCases) RemoveInvited(ctx context.Context, task primitive.ObjectID, ids []primitive.ObjectID) (primitive.ObjectID, error) {
	valid, err := u.repos.Person.ValidateUUID(ctx, ids)

	if err != nil {
		return primitive.NilObjectID, err
	}

	if !valid {
		return primitive.NilObjectID, errors.New("one or more invited person IDs do not exist")
	}

	u.repos.Task.RemoveInvited(ctx, task, ids)
	return task, nil
}
