package main

import (
	"fmt"
	"log"
	"net/http"
	"user"
)

func main() {
	fmt.Println("============== Go grocery Server Starting ===================")
	user.SetUpUserRoute()
}

func handleRoot(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		log.Println("In valid method")
		responseWriter.Write([]byte("Not a valid request"))
		http.Error(responseWriter, "Bad request, this route needs http post request...", http.StatusBadRequest)
	}
}
