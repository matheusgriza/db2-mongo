package repositories

import (
	"context"
	"task-api/internal/models"
	"task-api/internal/repositories/persons"
	"task-api/internal/repositories/tasks"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repositories struct {
	Person interface {
		GetPerson(ctx context.Context, id primitive.ObjectID) (*models.Person, error)
		GetAllPerson(ctx context.Context) ([]models.Person, error)
		AddPerson(ctx context.Context, newPerson models.Person) error
		ValidateUUID(ctx context.Context, ids []primitive.ObjectID) (bool, error)
	}

	Task interface {
		GetTask(ctx context.Context, id primitive.ObjectID) (*models.TaskGet, error)
		GetAllTask(ctx context.Context) ([]models.TaskGet, error)
		DeleteTask(ctx context.Context, id primitive.ObjectID) (*models.Task, error)
		AddTask(ctx context.Context, newTask models.Task) error
		UpdateTask(ctx context.Context, id primitive.ObjectID, task models.UpdateTaskRequest) error
		AddInvited(ctx context.Context, taskId primitive.ObjectID, personIds []primitive.ObjectID) error
		RemoveInvited(ctx context.Context, taskId primitive.ObjectID, personId []primitive.ObjectID) error
	}
}

func New(db *mongo.Database) *Repositories {
	return &Repositories{
		Person: persons.New(db),
		Task:   tasks.New(db),
	}
}
