package routes

import (
	handler "Users/Handler"

	"github.com/gorilla/mux"
)

func SetupRoutes()*mux.Router{
	r:=mux.NewRouter()
	r.HandleFunc("/items",handler.CreateItem).Methods("POST")
	r.HandleFunc("/items{id}",handler.GetItem).Methods("GET")
	r.HandleFunc("/items{id}",handler.UpdateItem).Methods("PUT")
	r.HandleFunc("/items{id}",handler.DeleteItem).Methods("DELETE")
	return r
}