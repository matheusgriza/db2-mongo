package tasks

import (
	"context"
	"task-api/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (t *Tasks) GetTask(ctx context.Context, id primitive.ObjectID) (*models.TaskGet, error) {
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.D{{Key: "_id", Value: id}}}},
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "persons"},
			{Key: "localField", Value: "Invited"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "Invited"},
		}}}}

	cursor, err := t.col.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if !cursor.Next(ctx) {
		return nil, mongo.ErrNoDocuments
	}

	var task models.TaskGet

	if err := cursor.Decode(&task); err != nil {
		return nil, err
	}
	return &task, nil
}

func (t *Tasks) DeleteTask(ctx context.Context, id primitive.ObjectID) (*models.Task, error) {
	var task models.Task
	err := t.col.FindOneAndDelete(ctx, bson.M{"_id": id}).Decode(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (t Tasks) GetAllTask(ctx context.Context) ([]models.TaskGet, error) {
	pipeline := mongo.Pipeline{
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "persons"},
			{Key: "localField", Value: "Invited"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "Invited"},
		}}}}

	cursor, err := t.col.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if !cursor.Next(ctx) {
		return nil, mongo.ErrNoDocuments
	}

	var task []models.TaskGet
	if err := cursor.All(ctx, &task); err != nil {
		return nil, err
	}
	return task, nil

}

func (t *Tasks) AddTask(ctx context.Context, newTask models.Task) error {
	_, err := t.col.InsertOne(ctx, newTask)

	if err != nil {
		return err
	}

	return nil
}

func (t *Tasks) UpdateTask(ctx context.Context, id primitive.ObjectID, task models.UpdateTaskRequest) error {

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

func (t *Tasks) AddInvited(ctx context.Context, taskId primitive.ObjectID, personIds []primitive.ObjectID) error {
	_, err := t.col.UpdateByID(
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
	return err
}

func (t *Tasks) RemoveInvited(ctx context.Context, taskId primitive.ObjectID, personIds []primitive.ObjectID) error {
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
