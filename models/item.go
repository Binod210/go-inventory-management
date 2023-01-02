package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductDecode struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Qty  int32  `json:"qty,omitempty"`
	Rate int32  `json:"rate,omitempty"`
}

type Product struct {
	Id   primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Name string             `json:"name,omitempty" bson:"name"`
	Qty  int32              `json:"qty,omitempty" bson:"qty"`
	Rate int32              `json:"rate,omitempty" bson:"rate"`
}
