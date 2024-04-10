package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectDB() {
	clientOptions := options.Client().ApplyURI("")
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
}

func GetMongoClient() *mongo.Client {
	return client
}
func GetContext() context.Context {
	return context.Background()

}
func FindAll(collection *mongo.Collection) ([]interface{}, error) {
	var results []interface{}
	cursor, err := collection.Find(GetContext().bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	err = cursor.All(GetContext(), &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}
func FindOneById(collection *mongo.Collection, filter interface{}) (interface{}, error) {
	var result interface{}
	err := collection.FindOne(GetContext(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func InsertOne(collection *mongo.Collection, document interface{}) error {
	_, err := collection.InsertOne(GetContext(), document)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return err

}
func UpdateOne(collection *mongo.Collection, filter interface{}, update interface{}) {
	_, err := collection.UpdateOne(GetContext(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
}
func DeleteOne(collection *mongo.Collection, filter interface{}) {
	_, err := collection.DeleteOne(GetContext(), filter)
	if err != nil {
		log.Fatal(err)
	}

}
