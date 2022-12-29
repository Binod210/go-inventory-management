package services

import (
	"net/http"

	"github.com/binod210/go-inventory-management/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	Collection *db.UserRepository
}

func NewUserService(Db *mongo.Database) UserService {
	collection := db.NewUserRepository(Db)
	return UserServiceImpl{
		Collection: &collection,
	}
}

func (us UserServiceImpl) CreateUser(w http.ResponseWriter, r *http.Request) {

}
