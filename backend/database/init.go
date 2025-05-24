package database

import (
    "context"
    "os"
    "fmt"
    "log"
    "time"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "github.com/joho/godotenv"
)

func ConnectDB() *mongo.Client {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Failed to load .env file")
    }
    clientOptions := options.Client().
        ApplyURI(os.Getenv("DATABASE_URI"))

    client, err := mongo.NewClient(clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to MongoDB")
    return client
}

var DB *mongo.Client = ConnectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
    return client.Database("Go-practice").Collection(collectionName)
}
