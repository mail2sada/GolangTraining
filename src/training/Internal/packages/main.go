package main

import (
	"fmt"

	_ "github.com/mail2sada/testpackage/math/advanced"

	"github.com/mail2sada/testpackage/math"

	mt "github.com/mail2sada/userpackage/math"
)

func init() {

	fmt.Println("First init")
}

func init() {

	fmt.Println("Second init")
}

func init() {

	fmt.Println("third init")
}

func init() {

	fmt.Println("Fourth init")
}
func main() {

	fmt.Println("Demo packages....", a)

	fmt.Println(math.B)

	math.TestPrint()
	math.TestMyPrints()

	fmt.Println(mt.Add(10, 20))
	fmt.Println(mt.Subtract(30, 10))
}
