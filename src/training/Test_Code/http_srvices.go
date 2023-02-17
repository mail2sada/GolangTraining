package main

import (
	"fmt"
	"net/http"
)

func handleRootAuth(res http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(res, "Hi From handleRootAuth")

}
func handleRootDB(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hi from handleRootDB")
}

type Server struct {
	service http.Server
}

func main() {

	fmt.Println("Demo: Http services....")

	handlerAuth := http.NewServeMux()
	handlerAuth.HandleFunc("/", handleRootAuth)

	handlerDB := http.NewServeMux()

	dbService := &http.Server{Addr: "127.0.0.1:2233", Handler: handlerDB}

	authService := &http.Server{Addr: "127.0.0.1:8080", Handler: handlerAuth}

	go func(mux *http.ServeMux, server *http.Server) {

		server.ListenAndServe()
	}(handlerDB, dbService)

	authService.ListenAndServe()

	http.ListenAndServe()

	authService.ListenAndServe()

}
