package services

import (
	"net/http"

	"github.com/binod210/go-inventory-management/authentication"
	"github.com/binod210/go-inventory-management/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	CreateUser(http.ResponseWriter, *http.Request)
	Login(http.ResponseWriter, *http.Request)
	UpdateUser(http.ResponseWriter, *http.Request)
	DeleteUser(http.ResponseWriter, *http.Request)
}

func NewUserService(Db *mongo.Database, auth *authentication.JWT) UserService {
	collection := db.NewUserRepository(Db)
	return &UserServiceImpl{
		Collection: collection,
		Auth:       auth,
	}
}
