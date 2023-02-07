package user

type User struct {
	username string
	address  string
}

func GetAllUsers() []User {

	res := []User{{username: "gyan", address: "Marathalli"}, {username: "Guru", address: "Hebbal"}}

	return res
}
