package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

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

//USER PUT request
func UpdateUser(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	var new_user User
	json.NewDecoder(req.Body).Decode(&new_user)
	err := UpdateUserInfo(params["uuid"], new_user)
	if err.Status != "" {
		json.NewEncoder(writer).Encode(err)
	}
}

//USER DELETE request
func DeleteUser(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	err := DeleteUserInfo(params["uuid"])
	if err.Status != "" {
		json.NewEncoder(writer).Encode(err)
	}
}
