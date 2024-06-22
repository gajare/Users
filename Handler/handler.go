package handler

import (
	models "Users/Models"
	service "Users/Service"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	err := service.CreateItem(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	item, err := service.GetItem(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(item)

}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var item models.Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	err := service.UpdateItem(params["id"], item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(item)

}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := service.DeleteItem(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
