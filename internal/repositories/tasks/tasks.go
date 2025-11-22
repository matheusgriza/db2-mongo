package tasks

import (
	"context"
	"fmt"
	"task-api/internal/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Tasks struct {
	col *mongo.Collection
}

func New(db *mongo.Database) *Tasks {
	return &Tasks{
		col: db.Collection("tasks"),
	}
}

func (t *Tasks) GetTask(ctx context.Context, id uuid.UUID) (*models.Task, error) {
	var task models.Task
	err := t.col.FindOne(ctx, bson.M{"Id": id}).Decode(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (t *Tasks) DeleteTask(ctx context.Context, id uuid.UUID) (*models.Task, error) {
	var task models.Task
	err := t.col.FindOneAndDelete(ctx, bson.M{"Id": id}).Decode(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (t Tasks) GetAllTask(ctx context.Context) ([]models.Task, error) {
	var tasks []models.Task

	cursor, err := t.col.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (t *Tasks) AddTask(ctx context.Context, newTask models.Task) error {
	_, err := t.col.InsertOne(ctx, newTask)

	if err != nil {
		return err
	}

	return nil
}

func (t *Tasks) UpdateTask(ctx context.Context, id uuid.UUID, task models.UpdateTaskRequest) error {

	_, err := t.col.UpdateByID(ctx,
		id,
		bson.M{
			"$set": task,
		})

	if err != nil {
		return err
	}
	return nil
}

func (t *Tasks) AddInvited(ctx context.Context, taskId uuid.UUID, personIds []uuid.UUID) error {
	res, err := t.col.UpdateOne(
		ctx,
		taskId,
		bson.M{
			"$addToSet": bson.M{
				"Invited": bson.M{
					"$each": personIds,
				},
			},
		},
	)
	if err != nil {
		return err
	}

	fmt.Println("Matched:", res.MatchedCount)
	fmt.Println("Modified:", res.ModifiedCount)

	return err
}

func (t *Tasks) RemoveInvited(ctx context.Context, taskId uuid.UUID, personIds []uuid.UUID) error {
	_, err := t.col.UpdateByID(
		ctx,
		taskId,
		bson.M{
			"$pull": bson.M{
				"Invited": bson.M{
					"$in": personIds,
				},
			},
		},
	)

	return err
}
