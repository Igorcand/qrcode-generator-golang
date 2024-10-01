package repository

import (
    "context"
    "qrcode-generator/internal/core/domain/qrcode"
    "go.mongodb.org/mongo-driver/mongo"
)

type MongoQRCodeRepository struct {
    collection *mongo.Collection
}

// NewMongoQRCodeRepository cria uma nova instância do repositório de QR codes
func NewMongoQRCodeRepository(client *mongo.Client, dbName, collectionName string) *MongoQRCodeRepository {
    return &MongoQRCodeRepository{
        collection: client.Database(dbName).Collection(collectionName),
    }
}

// Save salva um QR code no MongoDB
func (repo *MongoQRCodeRepository) Save(qrCode *qrcode.QRCode) error {
    _, err := repo.collection.InsertOne(context.TODO(), qrCode)
    return err
}

// FindByID busca um QR code por ID no MongoDB
func (repo *MongoQRCodeRepository) FindByID(id string) (*qrcode.QRCode, error) {
    var qrCode qrcode.QRCode
    err := repo.collection.FindOne(context.TODO(), map[string]string{"id": id}).Decode(&qrCode)
    return &qrCode, err
}
