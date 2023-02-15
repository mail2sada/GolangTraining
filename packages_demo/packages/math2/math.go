package math

import (
	"fmt"

	advanced "github.com/mail2sada/mavenir/math2/Advanced"
)

func init() {

	fmt.Println("math: init from math")
}

func Add(a, b int) {

	c := a + b

	fmt.Println(c)
}

func Sub(a, b int) {

	c := a - b

	fmt.Println(c)
}

func Square(a int) {
	advanced.Square(a)
}
