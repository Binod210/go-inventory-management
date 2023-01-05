package services

import (
	"net/http"

	"github.com/binod210/go-inventory-management/authentication"
	"github.com/binod210/go-inventory-management/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderService interface {
	SaveOrder(http.ResponseWriter, *http.Request)
	UpdateOrder(http.ResponseWriter, *http.Request)
	DeleteOrder(http.ResponseWriter, *http.Request)
	GetOrderDetail(http.ResponseWriter, *http.Request)
	GetOrders(http.ResponseWriter, *http.Request)
	GetOrdersByProductId(http.ResponseWriter, *http.Request)
}

func NewOrderService(Db *mongo.Database, auth *authentication.JWT) OrderService {
	repo := db.NewOrderRepository(Db)
	return &OrderServiceImpl{
		Repository: repo,
		Auth:       auth,
	}
}
