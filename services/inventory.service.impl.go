package services

import (
	"github.com/binod210/go-inventory-management/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type InventoryServiceImpl struct {
	Repository *db.InventoryRepository
}

func NewInventoryService(Db *mongo.Database) InventoryService {
	repo := db.NewInventoryRepository(Db)
	return &InventoryServiceImpl{
		Repository: &repo,
	}
}
