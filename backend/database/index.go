package database

import (
    "context"
    "log"
    "time"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func CreateIndexes(client *mongo.Client) {
    collection := GetCollection(DB, "urls")

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    indexModel := mongo.IndexModel{
        Keys: bson.D{{Key: "shortened_url", Value: 1}}, 
        Options: options.Index().SetUnique(true),
    }

    name, err := collection.Indexes().CreateOne(ctx, indexModel)
    if err != nil {
        log.Fatalf("Failed to create index: %v", err)
    }

    log.Println("Created index:", name)
}
