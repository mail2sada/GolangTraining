package main

import (
	"fmt"
	"gogrocery/user"
)

func main() {

	fmt.Println("GO GROCERY SERVER")
	var u user.User = user.User{}

	usr := user.GetAllUsers()

	fmt.Println(u, usr)
}
