package repository

import (
	"Birds/internal/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BirdRepository struct {
	db *mongo.Client
}

func (b *BirdRepository) Insert(ctx context.Context, bird models.Bird) error {
	_, err := b.db.Database("Birds").Collection("Birds").InsertOne(ctx, bird)
	return err
}

func (b *BirdRepository) Update(ctx context.Context, bird models.Bird) error {
	update := bson.M{"$set": bird}
	_, err := b.db.Database("Birds").Collection("Birds").UpdateOne(ctx, bson.M{"owner": bird.Owner}, update)
	return err
}

func (b *BirdRepository) FindByOwner(ctx context.Context, owner string) (models.Bird, error) {
	result := models.Bird{}
	return result, b.db.Database("Birds").Collection("Birds").FindOne(ctx, bson.M{"owner": owner}).Decode(&result)
}

func (b *BirdRepository) DeleteByOwner(ctx context.Context, owner string) error {
	_, err := b.db.Database("Birds").Collection("Birds").DeleteOne(ctx, bson.M{"owner": owner})
	return err
}

func NewBirdRepository(db *mongo.Client) *BirdRepository {
	return &BirdRepository{db: db}
}
