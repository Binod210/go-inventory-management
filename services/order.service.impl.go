package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/binod210/go-inventory-management/authentication"
	"github.com/binod210/go-inventory-management/db"
	"github.com/binod210/go-inventory-management/models"
	"github.com/gorilla/mux"
)

type OrderServiceImpl struct {
	Repository db.OrderRepository
	Auth       *authentication.JWT
}

func (os OrderServiceImpl) SaveOrder(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	err := os.Auth.VerifyToken(authHeader)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err)
		return
	}

	var order models.OrderDecode
	err = json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		log.Println("Error in Save order in OrderService", err)
		json.NewEncoder(w).Encode("cannot decode the json input")
		return
	}
	saveorder, err := os.Repository.Save(&order)
	if err != nil {
		log.Println("Error during saving to repository ", err)
		json.NewEncoder(w).Encode("Error during saving to repository")
		return
	}
	json.NewEncoder(w).Encode(saveorder)

}

func (os OrderServiceImpl) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	err := os.Auth.VerifyToken(authHeader)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err)
		return
	}

	var order models.OrderDecode
	err = json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		log.Println("Error in update order in OrderService", err)
		json.NewEncoder(w).Encode("cannot decode the json input")
		return
	}

	dbOrder, err := os.Repository.FindById(order.Id)
	if err != nil {
		log.Println("Error in fetching the order with order id ", err)
		json.NewEncoder(w).Encode(fmt.Sprintf("Cannot find  order with order id %v", order.Id))
		return
	}
	if order.Qty != 0 {
		dbOrder.Qty = order.Qty
	}
	saveOrder, err := os.Repository.Save(dbOrder)
	if err != nil {
		log.Println("error in saving update order ", err)
		json.NewEncoder(w).Encode("Could not update order")
		return
	}
	json.NewEncoder(w).Encode(saveOrder)

}

func (os OrderServiceImpl) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	err := os.Auth.VerifyToken(authHeader)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err)
		return
	}

	params := mux.Vars(r)
	id := params["id"]
	err = os.Repository.Delete(id)
	if err != nil {
		log.Println("error in deletion in element", err)
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode("Successfully deleted order")

}

func (os OrderServiceImpl) GetOrderDetail(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	err := os.Auth.VerifyToken(authHeader)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err)
		return
	}
}

func (os OrderServiceImpl) GetOrders(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	err := os.Auth.VerifyToken(authHeader)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err)
		return
	}

}

func (os OrderServiceImpl) GetOrdersByProductId(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	err := os.Auth.VerifyToken(authHeader)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err)
		return
	}

}
