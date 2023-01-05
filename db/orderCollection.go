package db

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/binod210/go-inventory-management/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderCollection struct {
	Collection *mongo.Collection
}

func NewOrderRepository(db *mongo.Database) OrderRepository {
	coll := db.Collection("Orders")
	return &OrderCollection{
		Collection: coll,
	}
}

func (oc *OrderCollection) Save(order *models.OrderDecode) (*models.OrderDecode, error) {

	if order == nil {
		return nil, errors.New("order cannot be null")
	}
	var isUpdate = false
	if order.Id != "" {
		isUpdate = true
	}
	saveOrder := convertToOrder(order)
	if isUpdate {
		filter := bson.M{"_id": saveOrder.Id}
		updateBson, err := convertToUpdateBson(saveOrder)
		if err != nil {
			return order, err
		}
		result, err := oc.Collection.UpdateOne(context.TODO(), filter, updateBson)
		if err != nil {
			return order, err
		}
		if result.MatchedCount == 0 && result.ModifiedCount == 0 {
			return order, errors.New("cannot update user")
		}
	} else {
		result, err := oc.Collection.InsertOne(context.TODO(), saveOrder)
		if err != nil {
			return nil, err
		}
		if len(result.InsertedID.(primitive.ObjectID).Hex()) == 0 {
			return order, errors.New("cannot save order")
		}
		return convertToOrderDecode(saveOrder), nil
	}
	return nil, nil
}

func (oc *OrderCollection) FindAll() ([]*models.OrderDecode, error) {
	cur, err := oc.Collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	var orders []*models.OrderDecode
	for cur.Next(context.TODO()) {
		order := models.Order{}
		err = cur.Decode(&order)
		if err != nil {
			log.Println(err)
			continue
		}
		orders = append(orders, convertToOrderDecode(&order))
	}
	return orders, nil

}

func (oc *OrderCollection) Delete(id string) error {
	objectId, _ := primitive.ObjectIDFromHex(id)
	result, err := oc.Collection.DeleteOne(context.TODO(), bson.M{"_id": objectId})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("order not found to delete")
	}
	return nil
}

func (oc *OrderCollection) FindById(id string) (*models.OrderDecode, error) {
	if id != "" {
		return nil, errors.New("id cannot be null")
	}

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid Id")
	}
	result := oc.Collection.FindOne(context.TODO(), bson.M{"_id": objectId})
	if result.Err() != nil {
		return nil, result.Err()
	}

	var order models.Order
	err = result.Decode(&order)
	if err != nil {
		return nil, err
	}
	return convertToOrderDecode(&order), nil
}

func (oc OrderCollection) FindByProductId(id string) ([]*models.OrderDecode, error) {

	cur, err := oc.Collection.Find(context.TODO(), bson.M{"productId": id})
	if err != nil {
		return nil, err
	}
	var orders []*models.OrderDecode
	for cur.Next(context.TODO()) {
		order := models.Order{}
		err = cur.Decode(&order)
		if err != nil {
			log.Println(err)
			continue
		}
		orders = append(orders, convertToOrderDecode(&order))
	}
	return orders, nil
}

func convertToOrder(orderDecode *models.OrderDecode) *models.Order {
	var id primitive.ObjectID
	var err error
	if orderDecode.Id == "" {
		id = primitive.NewObjectID()
	} else {
		id, err = primitive.ObjectIDFromHex(orderDecode.Id)
		log.Println(err)
	}
	var timestamp primitive.Timestamp
	if orderDecode.Timestamp == 0 {
		timestamp = primitive.Timestamp{T: uint32(time.Now().Unix())}
	}
	return &models.Order{
		Id:        id,
		ProductId: orderDecode.ProductId,
		Qty:       orderDecode.Qty,
		Rate:      orderDecode.Rate,
		Timestamp: timestamp,
	}
}

func convertToOrderDecode(order *models.Order) *models.OrderDecode {
	timestamp := order.Timestamp.T
	return &models.OrderDecode{
		Id:        order.Id.Hex(),
		ProductId: order.ProductId,
		Qty:       order.Qty,
		Rate:      order.Rate,
		Timestamp: int64(timestamp),
	}
}
