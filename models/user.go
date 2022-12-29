package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Name     string             `json:"name,omitempty" bson:"name"`
	Email    string             `json:"email,omitempty" bson:"email"`
	Password string             `json:"password,omitempty" bson:"password"`
}

type UserDecode struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
