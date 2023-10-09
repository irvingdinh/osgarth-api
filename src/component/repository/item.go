package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/irvingdinh/osgarth-api/src/component/config"
	"github.com/irvingdinh/osgarth-api/src/component/model"
)

type ItemRepository interface {
	FindOneBySlug(ctx context.Context, slug string) (*model.Item, error)
}

func NewItemRepository(
	mongoClient *mongo.Client,
) ItemRepository {
	return &itemRepositoryImpl{
		mongoClient: mongoClient,
	}
}

type itemRepositoryImpl struct {
	mongoClient *mongo.Client
}

func (i *itemRepositoryImpl) FindOneBySlug(ctx context.Context, slug string) (*model.Item, error) {
	col := i.mongoClient.Database(config.GetDatabaseConfig().Name).Collection("items")

	filter := bson.D{{"slug", slug}}

	var result model.Item

	err := col.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
