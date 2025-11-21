package persons

import (
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ValidateUUID(ctx context.Context, personCol *mongo.Collection, ids []uuid.UUID) (bool, error) {
	if len(ids) == 0 {
		return true, nil
	}

	count, err := personCol.CountDocuments(ctx, bson.M{
		"_id": bson.M{"$in": ids},
	})

	if err != nil {
		return false, err
	}

	return count == int64(len(ids)), nil
}
