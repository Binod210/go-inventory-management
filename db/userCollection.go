package db

import (
	"github.com/binod210/go-inventory-management/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserCollection struct {
	Collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) UserRepository {
	collection := db.Collection("User")
	return &UserCollection{
		Collection: collection,
	}
}

func (uc *UserCollection) Save(user *models.User) (*models.User, error) {
	return nil, nil

}

func (uc *UserCollection) SaveAll(users *[]models.User) ([]*models.User, error) {
	return nil, nil
}

func (uc *UserCollection) FindAll() ([]*models.User, error) {
	return nil, nil
}

func (uc *UserCollection) FindById(Id string) (*models.User, error) {
	return nil, nil
}
