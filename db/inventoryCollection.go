package db

import (
	"github.com/binod210/go-inventory-management/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type InventoryCollection struct {
	Collection *mongo.Collection
}

func NewInventoryRepository(Db *mongo.Database) InventoryRepository {
	collection := Db.Collection("Inventory")
	return &InventoryCollection{
		Collection: collection,
	}
}

func (col *InventoryCollection) Save(item *models.Item) (*models.Item, error) {
	return nil, nil
}

func (col *InventoryCollection) SaveAll(items []*models.Item) ([]*models.Item, error) {
	return nil, nil
}

func (col *InventoryCollection) FindAll() ([]*models.Item, error) {
	return nil, nil
}

func (col *InventoryCollection) FindById(Id string) (*models.Item, error) {
	return nil, nil
}
