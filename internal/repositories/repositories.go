package repositories

import (
	"context"
	"task-api/internal/models"
	"task-api/internal/repositories/persons"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repositories struct {
	Person interface {
		GetPerson(ctx context.Context, id uuid.UUID) (*models.Person, error)
		GetAllPerson(ctx context.Context) ([]models.Person, error)
		AddPerson(ctx context.Context, newPerson models.Person) error
	}

	Task interface {
		GetTask(id uuid.UUID) (*models.Task, error)
		GetAllTask() ([]*models.Task, error)
		AddTask(newTask models.Task) error
		UpdateTask(id uuid.UUID) error
	}
}

func New(db *mongo.Database) *Repositories {
	return &Repositories{
		Person: persons.New(db),
	}
}
