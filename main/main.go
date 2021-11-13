package main

import (
	"main/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/boards", controller.Create).Methods("POST")
	router.HandleFunc("/boards", controller.ReadAll).Methods("GET")
	router.HandleFunc("/boards/{id:[0-9]+}", controller.ReadById).Methods("GET")
	router.HandleFunc("/boards/{id:[0-9]+}", controller.ModifyById).Methods("PATCH")
	router.HandleFunc("/boards/{id:[0-9]+}", controller.DeleteById).Methods("DELETE")

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
