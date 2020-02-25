package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	my_mongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Client my_mongo.Client

func init() {
	client, err := my_mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}

	// Create connect
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	Client = *client
}

func LoginProcess(login string, pass string) (LogErr, LogSucc) {
	users := Client.Database("test").Collection("users")
	var login_succ LogSucc
	var log_err LogErr
	var result User //buffer for FindOne() function

	filter := bson.D{{"login", login}, {"pass", pass}} //will find only documents with login and pass value
	err := users.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log_err = LogErr{Status: "error", Message: "User cannot be authenticated!"}
		return log_err, login_succ
	}

	login_succ = LogSucc{Status: "success", Message: "User found and generated token",
		Token: "eyJhbGciOi-RkOM8Hjc5DYNJuqyEy3gvy_IMjcu2w-hl2yHilvPNP_UK0ocUxaKdsD5oS5fV-TYlfH_k",
		Exp:   "2020-03-20", TokenMsg: "use the token to access the endpoints!", Login: result}

	return log_err, login_succ
}

func AddNewUser(new_user User) {
	users := Client.Database("test").Collection("users")

	_, err := users.InsertOne(context.TODO(), new_user)
	if err != nil {
		log.Fatal(err)
	}
}

func GetUserInfo(uuid string) (LogErr, User) {
	users := Client.Database("test").Collection("users")
	var result User //buffer for FindOne() function
	var log_err LogErr

	filter := bson.D{{"uuiduser", uuid}}
	err := users.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log_err = LogErr{Status: "error", Message: "User cannot be found"}
		return log_err, result
	}
	return log_err, result
}

func UpdateUserInfo(uuid string, new_data User) LogErr {
	users := Client.Database("test").Collection("users")
	var log_err LogErr

	filter := bson.D{{"uuiduser", uuid}}
	update := bson.M{"$set": bson.M{"avatar_image": new_data.Avatar_image,
		"avatar_type": new_data.Avatar_type, "name": new_data.Name,
		"surname": new_data.Surname, "datastart": new_data.Datastart, "login": new_data.Login,
		"pass": new_data.Pass}}

	_, err := users.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log_err = LogErr{Status: "error", Message: "User cannot be found"}
		return log_err
	}

	return log_err
}

func DeleteUserInfo(uuid string) LogErr {
	users := Client.Database("test").Collection("users")
	var log_err LogErr

	filter := bson.D{{"uuiduser", uuid}}

	_, err := users.DeleteOne(context.TODO(), filter)
	if err != nil {
		log_err = LogErr{Status: "error", Message: "User cannot be found"}
		return log_err
	}

	return log_err
}
