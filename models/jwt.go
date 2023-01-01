package models

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `jsong:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}
