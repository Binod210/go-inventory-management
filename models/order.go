package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	Id        primitive.ObjectID  `json:"id,omitempty" bson:"_id"`
	ProductId string              `json:"product_id,omitempty" bson:"productId"`
	Qty       int                 `json:"qty,omitempty" bson:"qty"`
	Rate      float32             `json:"rate,omitempty" bson:"rate"`
	Timestamp primitive.Timestamp `json:"timestamp,omitempty" bson:"timestamp"`
}

type OrderDecode struct {
	Id        string  `json:"id,omitempty"`
	ProductId string  `json:"product_id,omitempty"`
	Qty       int     `json:"qty,omitempty"`
	Rate      float32 `json:"rate,omitempty"`
	Timestamp int64   `json:"timestamp,omitempty"`
}
