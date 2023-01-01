package services

import "net/http"

type UserService interface {
	CreateUser(http.ResponseWriter, *http.Request)
	Login(http.ResponseWriter, *http.Request)
	UpdateUser(http.ResponseWriter, *http.Request)
	DeleteUser(http.ResponseWriter, *http.Request)
}
