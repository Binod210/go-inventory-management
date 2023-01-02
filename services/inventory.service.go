package services

import "net/http"

type InventoryService interface {
	AddProduct(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
}
