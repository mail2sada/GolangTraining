package main

import "fmt"

type author struct {
	int
	string
	float32
}

func main() {
	fmt.Println("... Demonstraiting Methods")

	var au author
	au = au.setValues(10, "Hello..", 3.142)
	au.print()
}
func (a author) print() {
	fmt.Printf("int :%d\nstring :%s\nfloat32 :%f\n", a.int, a.string, a.float32)
}

func (auth author) setValues(a int, b string, c float32) (ret author) {
	auth.int = a
	auth.string = b
	auth.float32 = c

	return auth
}
