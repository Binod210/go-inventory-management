package services

import "net/http"

type UserService interface {
	CreateUser(http.ResponseWriter, *http.Request)
}
