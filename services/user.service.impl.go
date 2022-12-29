package services

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/binod210/go-inventory-management/db"
	"github.com/binod210/go-inventory-management/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	Collection db.UserRepository
}

func NewUserService(Db *mongo.Database) UserService {
	collection := db.NewUserRepository(Db)
	return &UserServiceImpl{
		Collection: collection,
	}
}

func (us *UserServiceImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateUser called")
	var user models.UserDecode
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		json.NewEncoder(w).Encode("Cannot decode message") //Todo --- Error structure
		return
	}
	saveUser, err := us.Collection.Save(&user)
	if err != nil {
		json.NewEncoder(w).Encode(err) //Todo --- Error structure
		return
	}
	json.NewEncoder(w).Encode(saveUser)

}

func (us *UserServiceImpl) Login(w http.ResponseWriter, r *http.Request) {

}

func (us *UserServiceImpl) UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func (us *UserServiceImpl) DeleteUser(w http.ResponseWriter, r *http.Request) {

}
