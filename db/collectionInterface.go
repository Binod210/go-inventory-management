package db

import "github.com/binod210/go-inventory-management/models"

type InventoryRepository interface {
	Save(*models.Item) (*models.Item, error)
	SaveAll([]*models.Item) ([]*models.Item, error)
	FindAll() ([]*models.Item, error)
	FindById(string) (*models.Item, error)
}

type UserRepository interface {
	Save(*models.UserDecode) (*models.UserDecode, error)
	SaveAll(*[]models.UserDecode) ([]*models.UserDecode, error)
	FindAll() ([]*models.UserDecode, error)
	FindById(string) (*models.UserDecode, error)
	FindByEmailAndUser(string, string) (bool, error)
	DeleteUser(string) (bool, error)
}
