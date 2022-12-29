package db

import "github.com/binod210/go-inventory-management/models"

type InventoryRepository interface {
	Save(*models.Item) (*models.Item, error)
	SaveAll([]*models.Item) ([]*models.Item, error)
	FindAll() ([]*models.Item, error)
	FindById(string) (*models.Item, error)
}

type UserRepository interface {
	Save(*models.User) (*models.User, error)
	SaveAll(*[]models.User) ([]*models.User, error)
	FindAll() ([]*models.User, error)
	FindById(string) (*models.User, error)
}
