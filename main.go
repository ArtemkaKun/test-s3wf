package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const IS_SERVER_ONLINE bool = true //simulating server status for PING request

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/ping", Ping).Methods("POST")
	router.HandleFunc("/auth/{login}/{pass}", Auth).Methods("POST")
	router.HandleFunc("/user", AddUser).Methods("POST")
	router.HandleFunc("/order", AddOrder).Methods("POST")
	router.HandleFunc("/user/{uuid}", DeleteUser).Methods("DELETE")
	router.HandleFunc("/user/{uuid}", UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{uuid}", GetUser).Methods("GET")
	router.HandleFunc("/order/{orderid}", GetOrder).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router)) //run server on 8000 port
}
