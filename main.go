package main

import (
	routes "Users/Routes"
	"log"
)

func main(){
	r:=routes.SetupRoutes()
	log.Fatal(":8080",r)
}