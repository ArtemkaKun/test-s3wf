package main

import (
	"encoding/json"
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
	router.HandleFunc("/user/{uuid}", GetUser).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router)) //run server on 8000 port

}

//PING request
func Ping(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(IS_SERVER_ONLINE)
}

//AUTH request
func Auth(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	err, login := LoginProcess(params["login"], params["pass"])
	if err.Status != "" {
		json.NewEncoder(writer).Encode(err)
	} else {
		json.NewEncoder(writer).Encode(login)
	}
}

//USER POST request
func AddUser(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var new_user User
	json.NewDecoder(req.Body).Decode(&new_user)
	AddNewUser(new_user)
}

//USER GET request
func GetUser(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	err, user := GetUserInfo(params["uuid"])
	if err.Status != "" {
		json.NewEncoder(writer).Encode(err)
	} else {
		json.NewEncoder(writer).Encode(user)
	}
}
