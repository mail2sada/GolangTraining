package main

import (
	"fmt"
	"net/http"
)

func handleRoot(res http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case "POST":
		///sdkslkdjlsd
	case "GET":
		//kdjslkjd
	case "PUT":
	case "DELETE":
	default:
	}

	fmt.Fprintln(res, "Welcome to my web server...")
}

func handleHi(res http.ResponseWriter, req *http.Request) {

	fmt.Fprintln(res, "Hello, I am good")

}

func handleHello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hi, I am good. How are you?")
}

func main() {

	fmt.Println("Demo Web server...")

	http.HandleFunc("/", handleRoot)

	http.HandleFunc("/hi", handleHi)

	http.HandleFunc("/hello", handleHello)

	http.ListenAndServe(":2233", nil)
}
