package user

import (
	"fmt"
	"net/http"
)

/*
this file will deal with creation deletion, update and authentication of user.
*/


func SetUpUserRoute(){
	http.HandleFunc("/user", handleGetUser) //this is for get
	http.HandleFunc("/user/create/", handleCreateUser)
	http.HandleFunc("/user/update", handleUpdateUser)
	http.HandleFunc("/user/delete", handleDeleteUser)
}

func handleGetUser(responseWriter http.ResponseWriter, request * http.Request) {

	http.Error(responseWriter, "no users found,", http.StatusOK)

}

func handleCreateUser(responseWriter http.ResponseWriter, request *http.Request) {

}

func handleUpdateUser (responseWriter http.ResponseWriter, request *http.Request) {

}

func handleDeleteUser (responseWriter http.ResponseWriter, request *http.Request) {

}



func AllIsWelle() {
	fmt.Println("User....")
}