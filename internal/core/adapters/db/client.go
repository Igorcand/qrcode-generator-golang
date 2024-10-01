package db

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "log"
    "time"
)

func NewMongoClient(uri string) (*mongo.Client, error) {
    clientOptions := options.Client().ApplyURI(uri)
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        return nil, err
    }

    if err = client.Ping(ctx, nil); err != nil {
        return nil, err
    }

    log.Println("Conexão com MongoDB estabelecida com sucesso!")
    return client, nil
}
