package db

import (
	"context"
	"errors"
	"log"

	"github.com/binod210/go-inventory-management/models"
	"go.mongodb.org/mongo-driver/bson"
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
	isupdate := false
	if user.Id != "" {
		isupdate = true
	}
	saveUser := convertToUser(user)
	if isupdate {
		filter := bson.M{"_id": saveUser.Id}
		updateBson, err := convertToUpdateBson(saveUser)
		if err != nil {
			return user, err
		}
		result, err := uc.Collection.UpdateOne(context.TODO(), filter, updateBson)
		if err != nil {
			return user, err
		}
		if result.MatchedCount == 0 && result.ModifiedCount == 0 {
			return user, errors.New("cannot update user")
		}
	} else {
		result, err := uc.Collection.InsertOne(context.TODO(), saveUser)
		if err != nil {
			return user, errors.New("cannot save user")
		}

		if len(result.InsertedID.(primitive.ObjectID).Hex()) == 0 {
			return user, errors.New("cannot save user")
		}
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
	if Id == "" {
		return nil, errors.New("id cannot be null")
	}
	objectId, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return nil, errors.New("invalid Id")
	}

	result := uc.Collection.FindOne(context.TODO(), bson.M{"_id": objectId})
	if result.Err() != nil {
		return nil, result.Err()
	}
	var user models.User
	err = result.Decode(&user)
	if err != nil {
		return nil, err
	}

	return convertToUserDecode(&user), nil
}

func (uc *UserCollection) FindByEmailAndUser(email string, password string) (bool, error) {
	log.Println("uc.FindByEmailAndUser()")
	if email == "" || password == "" {
		return false, errors.New("email or password if empty")
	}
	result := uc.Collection.FindOne(context.TODO(), bson.M{"email": email, "password": password})
	if result.Err() != nil {
		return false, result.Err()
	}
	var user models.User
	result.Decode(&user)
	if len(user.Id.Hex()) == 0 {
		return false, errors.New("wrong email or password")
	}
	return true, nil
}

func (uc *UserCollection) DeleteUser(id string) (bool, error) {
	objectId, _ := primitive.ObjectIDFromHex(id)
	result, err := uc.Collection.DeleteOne(context.TODO(), bson.M{"_id": objectId})
	if err != nil {
		return false, err
	}
	if result.DeletedCount == 0 {
		return false, nil
	}
	return true, nil
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
