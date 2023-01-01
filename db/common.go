package db

import "go.mongodb.org/mongo-driver/bson"

func convertToUpdateBson(data interface{}) (bson.D, error) {
	dataByte, err := bson.Marshal(data)
	if err != nil {
		return nil, err
	}
	var bsonData bson.M
	err = bson.Unmarshal(dataByte, &bsonData)
	if err != nil {
		return nil, err
	}
	return bson.D{{Key: "$set", Value: bsonData}}, nil
}
