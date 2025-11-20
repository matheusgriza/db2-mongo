package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	uri := "mongodb://root:1234@127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.5.9"
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected w MongoDB")

	collection := client.Database("testedb").Collection("dummy")

	_, err = collection.InsertOne(context.Background(), map[string]any{
		"name": "Matheus Doe",
		"age":  30,
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted Document")

	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())

	var users []bson.M

	if (err) = cursor.All(context.Background(), &users); err != nil {
		log.Fatal(err)
	}

	fmt.Print(users)

}
