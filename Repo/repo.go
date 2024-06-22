package repo

import (
	models "Users/Models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"	
)

var collection *mongo.Collection

func init() {
	clientoptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientoptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("my_db").Collection("my_collection")
	//fmt.Println("collection :", collection)
}

func InsertItem(item models.Item) error {
	_, err := collection.InsertOne(context.Background(), item)
	return err
}

func FindItemByID(id primitive.ObjectID) (models.Item, error) {
	var item models.Item
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&item)
	return item, err
}

func UpdateItem(id primitive.ObjectID, item models.Item) error {
	_, err := collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": item})
	return err
}

func DeleteItem(id primitive.ObjectID) error {
	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}
