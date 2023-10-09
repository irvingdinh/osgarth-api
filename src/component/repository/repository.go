package repository

import "go.mongodb.org/mongo-driver/mongo"

type Repository interface {
	ItemRepository() ItemRepository
}

func New(
	mongoClient *mongo.Client,
) Repository {
	return &repositoryImpl{
		itemRepository: NewItemRepository(mongoClient),
	}
}

type repositoryImpl struct {
	itemRepository ItemRepository
}

func (i *repositoryImpl) ItemRepository() ItemRepository {
	return i.itemRepository
}
