package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	customer "github.com/mail2sada/OnlineShop/Users"
)

func init() {

	log.SetPrefix("[OnlineShop]")
	log.SetFlags(log.Ldate | log.Lshortfile)

	log.Println("Welcome to Onle Shoppe")
}

func handleRootGet(res http.ResponseWriter, req *http.Request) {

}

func handleRootPost(res http.ResponseWriter, req *http.Request) {

}

func main() {

	fileserver := http.FileServer(http.Dir("./static"))
	myMux := mux.NewRouter()

	customer.InitUserRoutes(myMux)

	myMux.Handle("/", fileserver)

	http.ListenAndServe("127.0.0.1:3344", myMux)

}
