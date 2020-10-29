package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arfandidts/dts-be-pengenalan-microservice/menu-service/handler"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/add-menu", http.HandlerFunc(handler.AddMenu))

	fmt.Println("Menu service listen on port :5000")
	log.Panic(http.ListenAndServe(":5000", router))
}
