package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// var apiGateway *http.Server = nil

// var auth *http.Server = nil

type mystruct struct {
	Name string `json:"name"`

	Age int `jason:"age"`
}

func init() {

	log.SetPrefix("[HTTP_SERVER]")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("=============== HTTP Server Init Complete===========")
}

func handleRoot(response http.ResponseWriter, request *http.Request) {
	log.Printf("In handleRoot...")

	var bodystruct = mystruct{}

	response.Write([]byte("I am too busy today\n"))

	body, _ := ioutil.ReadAll(request.Body)

	auth := request.Header.Get("Authorization")
	log.Printf(request.Method)
	switch strings.ToLower(request.Method) {

	case "get":
		response.Write([]byte("You are in get"))
		err := json.Unmarshal(body, &bodystruct)
		log.Println(bodystruct.Name, bodystruct.Age, err)
		response.Write([]byte(auth))

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)

	case "post":
		response.Write([]byte("You are in post"))
	case "put":
		response.Write([]byte("You are in put"))
	case "delete":
		response.Write([]byte("You are in delete"))
	}
}

func handleHello(response http.ResponseWriter, request *http.Request) {

	log.Printf("In handleHello...")

	//keys, _ := request.URL.Query()["id"]

	// id := keys[0]

	// fmt.Println(id)
	id := request.FormValue("id")

	fmt.Println(id)

	ctx := request.Context()
	body := request.Body
	header := request.Header
	log.Println(ctx, "   ", body, "  ", header)
	response.Write([]byte("Hello its fantastic day today..."))
}

func handleHowAreYou(response http.ResponseWriter, request *http.Request) {
	log.Printf("In handleHowAreYou...")

	response.Write([]byte("I am really doing good, how are you?"))
}

func main() {

	log.Println("Startig http server....")

	log.Println("Setting up handler...")

	fileserver := http.FileServer(http.Dir("./data"))

	http.Handle("/", fileserver)

	http.HandleFunc("/hello", handleHello)

	http.HandleFunc("/howareyou", handleHowAreYou)
	log.Println("handler setting complete...")

	http.ListenAndServe(":3333", nil)

}
