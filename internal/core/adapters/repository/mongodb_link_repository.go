package repository

import (
	"context"
	"qrcode-generator/internal/core/domain/link"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoLinkRepository struct{
	collection *mongo.Collection
}

func NewMongoLinkRepository(client *mongo.Client, dbName string, collectionName string) *MongoLinkRepository{
	return &MongoLinkRepository{
		collection: client.Database(dbName).Collection(collectionName),
	}
}

func (repo *MongoLinkRepository) Save(link *link.Link) error{
	_, err := repo.collection.InsertOne(context.TODO(),  link)
	return err
}

func (repo *MongoLinkRepository) FindByURL(url string) (*link.Link, error){
	var link link.Link
	err := repo.collection.FindOne(context.TODO(), map[string]string{"url": url}).Decode(&link)
	return &link, err
}