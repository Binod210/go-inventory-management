package db

import "github.com/binod210/go-inventory-management/models"

type InventoryRepository interface {
	Save(*models.ProductDecode) (*models.ProductDecode, error)
	SaveAll([]*models.ProductDecode) ([]*models.ProductDecode, error)
	FindAll() ([]*models.ProductDecode, error)
	FindById(string) (*models.ProductDecode, error)
	DeleteProduct(string) (bool, error)
}

type UserRepository interface {
	Save(*models.UserDecode) (*models.UserDecode, error)
	SaveAll(*[]models.UserDecode) ([]*models.UserDecode, error)
	FindAll() ([]*models.UserDecode, error)
	FindById(string) (*models.UserDecode, error)
	FindByEmailAndUser(string, string) (bool, error)
	DeleteUser(string) (bool, error)
}

type OrderRepository interface {
	Save(*models.OrderDecode) (*models.OrderDecode, error)
	FindAll() ([]*models.OrderDecode, error)
	Delete(string) error
	FindById(string) (*models.OrderDecode, error)
	FindByProductId(string) ([]*models.OrderDecode, error)
}
