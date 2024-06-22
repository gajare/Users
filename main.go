package main

import (
	routes "Users/Routes"
	"Users/middleware"
	"log"
	"net/http"
)

func main() {
	r := routes.SetupRoutes()
	//use middleware
	r.Use(middleware.LoggingMidleware)
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
