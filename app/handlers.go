package app

import (
	"github.com/binod210/go-inventory-management/services"
)

func (a *App) createHandlers() {
	userService := services.NewUserService(a.Database.DB)
	// inventoryService:= services.InventoryService(a.Database.DB)

	a.Router.HandleFunc("/user", userService.CreateUser).Methods("POST")

}
