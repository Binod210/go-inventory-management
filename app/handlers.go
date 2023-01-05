package app

import (
	"time"

	"github.com/binod210/go-inventory-management/authentication"
	"github.com/binod210/go-inventory-management/services"
)

func (a *App) createHandlers() {
	auth := authentication.NewJWT("my_secret", 120*time.Minute)
	userService := services.NewUserService(a.Database.DB, auth)
	inventoryService := services.NewInventoryService(a.Database.DB, auth)
	orderService := services.NewOrderService(a.Database.DB, auth)

	a.Router.HandleFunc("/user", userService.CreateUser).Methods("POST")
	a.Router.HandleFunc("/user/login", userService.Login).Methods("POST")
	a.Router.HandleFunc("/user", userService.UpdateUser).Methods("PUT")
	a.Router.HandleFunc("/user/{id}", userService.DeleteUser).Methods("DELETE")

	a.Router.HandleFunc("/inventory", inventoryService.AddProduct).Methods("POST")
	a.Router.HandleFunc("/inventory", inventoryService.AddProduct).Methods("PUT")
	a.Router.HandleFunc("/inventory", inventoryService.GetProducts).Methods("GET")
	a.Router.HandleFunc("/inventory/{id}", inventoryService.DeleteProduct).Methods("DELETE")

	a.Router.HandleFunc("/order", orderService.SaveOrder).Methods("POST")
	a.Router.HandleFunc("/order", orderService.UpdateOrder).Methods("PUT")
	a.Router.HandleFunc("/order/{id}", orderService.DeleteOrder).Methods("DELETE")
	a.Router.HandleFunc("/order/product/{productId}", orderService.GetOrdersByProductId).Methods("GET")
	a.Router.HandleFunc("/order/{id}", orderService.GetOrderDetail).Methods("GET")
	a.Router.HandleFunc("/order/all", orderService.GetOrders).Methods("GET")

}
