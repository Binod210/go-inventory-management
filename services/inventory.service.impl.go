package services

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/binod210/go-inventory-management/authentication"
	"github.com/binod210/go-inventory-management/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type InventoryServiceImpl struct {
	Repository *db.InventoryRepository
	Auth       *authentication.JWT
}

func NewInventoryService(Db *mongo.Database, auth *authentication.JWT) InventoryService {
	repo := db.NewInventoryRepository(Db)
	return &InventoryServiceImpl{
		Repository: &repo,
		Auth:       auth,
	}
}

func (is *InventoryServiceImpl) AddProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("Add Product called")
	authHeader := r.Header.Get("Authorization")
	err := is.Auth.VerifyToken(authHeader)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err)
		return
	}

}

func (is *InventoryServiceImpl) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Product called")

	authHeader := r.Header.Get("Authorization")
	err := is.Auth.VerifyToken(authHeader)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err)
		return
	}
}
