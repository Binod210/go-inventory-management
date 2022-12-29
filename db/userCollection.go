package db

import (
	"context"
	"errors"
	"log"

	"github.com/binod210/go-inventory-management/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (uc *UserCollection) Save(user *models.UserDecode) (*models.UserDecode, error) {
	log.Println("Save UserCollection called", user)
	if user == nil {
		return user, errors.New("user is null")
	}

	saveUser := convertToUser(user)
	result, err := uc.Collection.InsertOne(context.TODO(), saveUser)
	if err != nil {
		return user, errors.New("cannot save user")
	}

	if len(result.InsertedID.(primitive.ObjectID).Hex()) == 0 {
		return user, errors.New("cannot save user")
	}

	return convertToUserDecode(saveUser), nil
}

func (uc *UserCollection) SaveAll(users *[]models.UserDecode) ([]*models.UserDecode, error) {
	return nil, nil
}

func (uc *UserCollection) FindAll() ([]*models.UserDecode, error) {
	return nil, nil
}

func (uc *UserCollection) FindById(Id string) (*models.UserDecode, error) {
	return nil, nil
}

func convertToUser(userDecode *models.UserDecode) *models.User {
	var Id primitive.ObjectID
	var err error
	if userDecode.Id == "" {
		Id = primitive.NewObjectID()
	} else {
		Id, err = primitive.ObjectIDFromHex(userDecode.Id)
		log.Println(err)

	}
	return &models.User{
		Id:       Id,
		Name:     userDecode.Name,
		Email:    userDecode.Email,
		Password: userDecode.Password,
	}

}

func convertToUserDecode(user *models.User) *models.UserDecode {
	return &models.UserDecode{
		Id:    user.Id.Hex(),
		Name:  user.Name,
		Email: user.Email,
	}
}
