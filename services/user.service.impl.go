package services

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/binod210/go-inventory-management/authentication"
	"github.com/binod210/go-inventory-management/db"
	"github.com/binod210/go-inventory-management/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	Collection db.UserRepository
	Auth       *authentication.JWT
}

func NewUserService(Db *mongo.Database, auth *authentication.JWT) UserService {
	collection := db.NewUserRepository(Db)
	return &UserServiceImpl{
		Collection: collection,
		Auth:       auth,
	}
}

func (us *UserServiceImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateUser called")
	authHeader := r.Header.Get("Authorization")
	err := us.Auth.VerifyToken(authHeader)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err)
		return
	}
	var user models.UserDecode
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&user)
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
	var user models.UserDecode
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	userfind, err := us.Collection.FindByEmailAndUser(user.Email, user.Password)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	if !userfind {
		json.NewEncoder(w).Encode("Wrong username or password")
		return
	}
	token, err := us.Auth.GenerateToken(user.Email, "Admin")
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(token)

}

func (us *UserServiceImpl) UpdateUser(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	err := us.Auth.VerifyToken(authHeader)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err)
		return
	}

	var request models.UserDecode
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(err)
		return
	}

	user, err := us.Collection.FindById(request.Id)
	if err != nil {
		log.Println("Error in finding user ", user)
		json.NewEncoder(w).Encode(err)
		return
	}
	if request.Email != "" {
		user.Email = request.Email
	}

	if request.Name != "" {
		user.Name = request.Name
	}

	if request.Password != "" {
		user.Password = request.Password
	}
	user, err = us.Collection.Save(user)
	if err != nil {
		log.Println("Error in saving ", err)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(user)

}

func (us *UserServiceImpl) DeleteUser(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	err := us.Auth.VerifyToken(authHeader)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err)
		return
	}

	params := mux.Vars(r)
	id := params["id"]
	deleted, err := us.Collection.DeleteUser(id)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	if !deleted {
		json.NewEncoder(w).Encode("Failed to delete user")
		return
	}
	json.NewEncoder(w).Encode("Successfully deleted user")
}
