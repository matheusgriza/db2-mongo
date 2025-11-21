package persons

import (
	"context"
	"task-api/internal/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Persons struct {
	col *mongo.Collection
}

func New(db *mongo.Database) *Persons {
	return &Persons{col: db.Collection("persons")}
}

func (p *Persons) GetPerson(ctx context.Context, id uuid.UUID) (*models.Person, error) {
	var person models.Person
	err := p.col.FindOne(ctx, bson.M{"id": id}).Decode(&person)
	if err != nil {
		return nil, err
	}
	return &person, nil
}

func (p Persons) GetAllPerson(ctx context.Context) ([]models.Person, error) {
	var persons []models.Person

	cursor, err := p.col.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &persons); err != nil {
		return nil, err
	}

	return persons, nil
}

func (p *Persons) AddPerson(ctx context.Context, newPerson models.Person) error {
	_, err := p.col.InsertOne(ctx, newPerson)

	if err != nil {
		return err
	}

	return nil
}

func (p *Persons) ValidateUUID(ctx context.Context, ids []uuid.UUID) (bool, error) {
	count, err := p.col.CountDocuments(ctx, bson.M{
		"_id": bson.M{"$in": ids},
	})

	if err != nil {
		return false, err
	}

	return count == int64(len(ids)), nil
}
