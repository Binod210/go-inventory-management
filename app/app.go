package app

import (
	"log"
	"net/http"

	"github.com/binod210/go-inventory-management/db"
	"github.com/gorilla/mux"
)

type App struct {
	Router   *mux.Router
	Database db.DataBase
}

const (
	ConnString = "mongodb://localhost:27017"
)

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.Database = db.NewDatabase(ConnString)
	a.createHandlers()
}

func (a *App) Run(addr string) {
	err := http.ListenAndServe(addr, a.Router)
	if err != nil {
		log.Fatalf("Could not start Server %v", err)
	}
}
