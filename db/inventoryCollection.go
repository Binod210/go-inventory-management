package db

import (
	"context"
	"errors"
	"log"

	"github.com/binod210/go-inventory-management/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type InventoryCollection struct {
	Collection *mongo.Collection
}

func NewInventoryRepository(Db *mongo.Database) InventoryRepository {
	collection := Db.Collection("Inventory")
	return &InventoryCollection{
		Collection: collection,
	}
}

func (col *InventoryCollection) Save(product *models.ProductDecode) (*models.ProductDecode, error) {
	if product == nil {
		return nil, errors.New("product could not be nulls")
	}
	isUpdate := false
	if product.Id != "" {
		isUpdate = true
	}
	saveProduct := convertToProduct(product)
	if isUpdate {
		filter := bson.M{"_id": saveProduct.Id}
		updateBson, err := convertToUpdateBson(saveProduct)
		if err != nil {
			return product, err
		}
		result, err := col.Collection.UpdateOne(context.TODO(), filter, updateBson)
		if err != nil {
			return product, err
		}
		if result.MatchedCount == 0 && result.ModifiedCount == 0 {
			return product, errors.New("cannot update user")
		}
	} else {
		result, err := col.Collection.InsertOne(context.TODO(), saveProduct)
		if err != nil {
			return product, errors.New("cannot save user")
		}

		if len(result.InsertedID.(primitive.ObjectID).Hex()) == 0 {
			return product, errors.New("cannot save user")
		}

	}
	return convertToProductDecode(saveProduct), nil
}

func (col *InventoryCollection) SaveAll(products []*models.ProductDecode) ([]*models.ProductDecode, error) {
	return nil, nil
}

func (col *InventoryCollection) FindAll() ([]*models.ProductDecode, error) {
	return nil, nil
}

func (col *InventoryCollection) FindById(Id string) (*models.ProductDecode, error) {
	if Id != "" {
		return nil, errors.New("id cannot be null")
	}

	objectId, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return nil, errors.New("invalid Id")
	}
	result := col.Collection.FindOne(context.TODO(), bson.M{"_id": objectId})
	if result.Err() != nil {
		return nil, result.Err()
	}

	var product models.Product
	err = result.Decode(&product)
	if err != nil {
		return nil, err
	}
	return convertToProductDecode(&product), nil
}

func (col *InventoryCollection) DeleteProduct(Id string) (bool, error) {
	objectId, _ := primitive.ObjectIDFromHex(Id)
	result, err := col.Collection.DeleteOne(context.TODO(), bson.M{"_id": objectId})
	if err != nil {
		return false, err
	}
	if result.DeletedCount == 0 {
		return false, nil
	}
	return true, nil
}

func convertToProduct(productDecode *models.ProductDecode) *models.Product {
	var id primitive.ObjectID
	var err error
	if productDecode.Id == "" {
		id = primitive.NewObjectID()
	} else {
		id, err = primitive.ObjectIDFromHex(productDecode.Id)
		log.Println(err)

	}
	return &models.Product{
		Id:   id,
		Name: productDecode.Name,
		Qty:  productDecode.Qty,
		Rate: productDecode.Rate,
	}

}

func convertToProductDecode(product *models.Product) *models.ProductDecode {
	return &models.ProductDecode{
		Id:   product.Id.Hex(),
		Name: product.Name,
		Qty:  product.Qty,
		Rate: product.Rate,
	}
}
