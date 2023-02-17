package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

/*
/ Welcome to our http server

/hello I am doing good how are you

/ hi Hi great to see you
*/

func handleRoot(res http.ResponseWriter, req *http.Request) {

	// res.Write([]byte("Welcome to our http server"))

	// fmt.Println("Method used for the call is ", req.Method)

	// parameters := req.URL.Query()

	// fmt.Println("Parameter is ", parameters)

	// header := req.Header

	// contentType := header.Get("Content-Type")

	// body := req.Body

	// fmt.Println(body)

	// fmt.Println(contentType)

	// switch strings.ToLower(req.Method) {
	// case "get":
	// 	//handle get
	// 	fmt.Println("In get")
	// case "post":
	// 	// handle post
	// 	fmt.Println("In post..")
	// }

	//contentType := req.Header.Get("Content-Type")

	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Content-Length", strconv.Itoa(len("We have successfully service request....")))
	//res.Write([]byte("We have successfully service request...."))
	fmt.Fprintf(res, "We have successfully service request....")
	res.WriteHeader(http.StatusOK)

}

func handleHello(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("hello I am doing good how are you"))

}

func handleHi(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hi great to see you"))
	fmt.Println("Method used for the call is ", req.Method)
}

func main() {

	fmt.Println("HTTP Server.....")

	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/hello", handleHello)
	http.HandleFunc("/hi", handleHi)

	err := http.ListenAndServe("127.0.0.1:1234", nil)

	if err != nil {

		log.Fatalln("Could not start server ", err)
	}

}
