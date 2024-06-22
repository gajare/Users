package service

import (
	models "Users/Models"
	repo "Users/Repo"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateItem(item models.Item) error {
	return repo.InsertItem(item)
}

func GetItem(id string) (models.Item, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	return repo.FindItemByID(objID)
}

func UpdateItem(id string, item models.Item)error{
	objID,_:=primitive.ObjectIDFromHex(id)
	return repo.UpdateItem(objID,item)
}

func DeleteItem(id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	return repo.DeleteItem(objID)
}
