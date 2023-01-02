package services

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/binod210/go-inventory-management/authentication"
	"github.com/binod210/go-inventory-management/db"
	"github.com/binod210/go-inventory-management/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type InventoryServiceImpl struct {
	Repository db.InventoryRepository
	Auth       *authentication.JWT
}

func NewInventoryService(Db *mongo.Database, auth *authentication.JWT) InventoryService {
	repo := db.NewInventoryRepository(Db)
	return &InventoryServiceImpl{
		Repository: repo,
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

	var product models.ProductDecode
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	savedProduct, err := is.Repository.Save(&product)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(savedProduct)

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

	var product models.ProductDecode
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	if product.Id == "" {
		json.NewEncoder(w).Encode("Id cannot be null")
		return
	}

	dbProduct, err := is.Repository.FindById(product.Id)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	if product.Name == "" {
		dbProduct.Name = product.Name
	}
	if product.Qty == 0 {
		dbProduct.Qty = product.Qty
	}
	if product.Rate == 0 {
		dbProduct.Rate = product.Rate
	}
	dbProduct, err = is.Repository.Save(dbProduct)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(dbProduct)
}
