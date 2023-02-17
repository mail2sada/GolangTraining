package customer

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	Id      int    `json:id`
	Name    string `json:name`
	Email   string `json:email`
	Address string `json:address`
}

var userList = []User{}

func init() {
	log.Println("User init complete")
}

func handleUserCreate(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/text")
	if req.Method != "POST" {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Bad request expected post"))
		return
	}

	if req.Header.Get("Content-Type") != "application/json" {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Bad request expected Content-Type is application/json"))
		return
	}

	defer req.Body.Close()

	bodyData, err := io.ReadAll(req.Body)

	fmt.Println(string(bodyData))

	if err != nil {
		res.WriteHeader(http.StatusNoContent)
		log.Println("Recevied error ", err)
		res.Write([]byte("No Content need json"))
	}

	var user User

	err = json.Unmarshal(bodyData, &user)

	if err != nil {
		log.Println("Recevied error ", err)
		res.WriteHeader(http.StatusNoContent)
		res.Write([]byte("No Content need json"))
	}

	userList = append(userList, user)

	res.Header().Set("Content-Type", "application/text")
	res.WriteHeader(http.StatusOK)
	log.Println("User created")

}

func handleGetUser(res http.ResponseWriter, req *http.Request) {

	log.Println("handleGetUser")

	id := mux.Vars(req)["id"]

	idInt, _ := strconv.Atoi(id)

	var data = []byte{}

	for _, val := range userList {

		if val.Id == idInt {
			data, _ = json.Marshal(val)
		}
	}

	if len(data) == 0 {
		res.WriteHeader(http.StatusNotFound)
	}

	res.Write(data)
	res.WriteHeader(http.StatusOK)

	fmt.Println(id)

}

func handleAllUser(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-tyoe", "application.josn")

	data, _ := json.Marshal(userList)

	res.Header().Set("Content-Length", strconv.Itoa(len(data)))

	res.WriteHeader(http.StatusOK)
	res.Write(data)

}

func handleDeleteUser(res http.ResponseWriter, req *http.Request) {

}

func handleUpdateUser(res http.ResponseWriter, req *http.Request) {

}

func InitUserRoutes(route *mux.Router) {

	userList = append(userList, User{Id: 1, Name: "Anil kumar", Email: "anil.kumar@abc.com", Address: "#123 BTM Layout"})
	userList = append(userList, User{Id: 2, Name: "Vishal kumar", Email: "vishal.kumar@abc.com", Address: "#456 BTM Layout"})

	fmt.Println("Intializing UserRoutes..")

	route.HandleFunc("/user", handleUserCreate).Methods("POST")
	route.HandleFunc("/user/{id}", handleGetUser).Methods("GET")
	route.HandleFunc("/users", handleAllUser).Methods("GET")
	route.HandleFunc("/user/{id}", handleDeleteUser).Methods("DELETE")
	route.HandleFunc("/user/{id}", handleUpdateUser).Methods("PUT")

}
