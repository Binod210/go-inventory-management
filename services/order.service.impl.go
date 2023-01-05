package services

import (
	"net/http"

	"github.com/binod210/go-inventory-management/authentication"
	"github.com/binod210/go-inventory-management/db"
)

type OrderServiceImpl struct {
	Repository db.OrderRepository
	Auth       *authentication.JWT
}

func (os OrderServiceImpl) SaveOrder(w http.ResponseWriter, r *http.Request) {

}

func (os OrderServiceImpl) UpdateOrder(w http.ResponseWriter, r *http.Request) {

}

func (os OrderServiceImpl) DeleteOrder(w http.ResponseWriter, r *http.Request) {

}

func (os OrderServiceImpl) GetOrderDetail(w http.ResponseWriter, r *http.Request) {

}

func (os OrderServiceImpl) GetOrders(w http.ResponseWriter, r *http.Request) {

}

func (os OrderServiceImpl) GetOrdersByProductId(w http.ResponseWriter, r *http.Request) {

}
