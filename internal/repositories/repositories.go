package repositories

import (
	"context"
	"task-api/internal/models"
	"task-api/internal/repositories/persons"
	"task-api/internal/repositories/tasks"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repositories struct {
	Person interface {
		GetPerson(ctx context.Context, id uuid.UUID) (*models.Person, error)
		GetAllPerson(ctx context.Context) ([]models.Person, error)
		AddPerson(ctx context.Context, newPerson models.Person) error
		ValidateUUID(ctx context.Context, ids []uuid.UUID) (bool, error)
	}

	Task interface {
		GetTask(ctx context.Context, id uuid.UUID) (*models.Task, error)
		GetAllTask(ctx context.Context) ([]models.Task, error)
		AddTask(ctx context.Context, newTask models.Task) error
		UpdateTask(ctx context.Context, id uuid.UUID, task models.UpdateTaskRequest) error
		AddInvited(ctx context.Context, taskId uuid.UUID, personIds []uuid.UUID) error
		RemoveInvited(ctx context.Context, taskId uuid.UUID, personId []uuid.UUID) error
	}
}

func New(db *mongo.Database) *Repositories {
	return &Repositories{
		Person: persons.New(db),
		Task:   tasks.New(db),
	}
}
