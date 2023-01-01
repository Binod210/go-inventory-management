package authentication

import (
	"fmt"
	"strings"
	"time"

	"github.com/binod210/go-inventory-management/models"
	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	Secret string
	Expiry time.Duration
}

func NewJWT(secret string, expiry time.Duration) *JWT {
	return &JWT{
		Secret: secret,
		Expiry: expiry,
	}
}

func (j *JWT) GenerateToken(email string, role string) (string, error) {
	expiryTime := time.Now().Add(j.Expiry)
	claims := models.Claims{
		Username: email,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiryTime.Unix(),
		},
	}

	jwtKey := []byte(j.Secret)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}

func (j *JWT) VerifyToken(authHeader string) error {
	secret := []byte(j.Secret)
	token := bearerAuthHeader(authHeader)
	if token == "" {
		return fmt.Errorf("invalid token")
	}

	claims := &models.Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return fmt.Errorf("JWT token verification error: %s", err.Error())

		}
		return fmt.Errorf("JWT token verification error: %s", err.Error())
	}

	if !tkn.Valid {
		return fmt.Errorf("not a valid token")
	}
	return nil
}

func bearerAuthHeader(authHeader string) string {
	if authHeader == "" {
		return ""
	}

	parts := strings.Split(authHeader, "Bearer")
	if len(parts) != 2 {
		return ""
	}

	token := strings.TrimSpace(parts[1])
	if len(token) < 1 {
		return ""
	}

	return token
}
