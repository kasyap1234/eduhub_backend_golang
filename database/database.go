package database

import (
	"context"
	"log"

	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var client *mongo.Client

func ConnectDB() *mongo.Client{
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client; 
	

}

func GetMongoClient() *mongo.Client {
	return client
}
func GetContext() context.Context {
	return context.Background()

}
func FindAll(collection *mongo.Collection) ([]interface{}, error) {
	var results []interface{}
	cursor, err := collection.Find(GetContext(), bson.M{})
	if err != nil {
		log.Printf("Error executing Find: %v", err)
		return nil, err
	}

	err = cursor.All(GetContext(), &results)
	if err != nil {
		log.Printf("Error retrieving documents: %v", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

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
func DeleteOne(collection *mongo.Collection, filter interface{}) error {
	_, err := collection.DeleteOne(GetContext(), filter)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return err
}
