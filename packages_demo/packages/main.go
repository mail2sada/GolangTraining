package main

import (
	"fmt"

	_ "github.com/mail2sada/mavenir/blindimport"
	math "github.com/mail2sada/mavenir/math2"
	advanced "github.com/mail2sada/mavenir/math2/Advanced"
	mt "github.com/mail2sada/mavenir/mavtime"
	tt "github.com/mail2sada/mavenir/mavtime/testtime"

	_ "math"
)

func init() {
	fmt.Println("main: Init Method ...")
}

func init() {
	fmt.Println("main: init from 2nd init method")
}

func main() {

	fmt.Println("Demo: Packages..")

	fmt.Println("Value of a is ", a)

	fmt.Println("Value of A from mavtime package is ", mt.A)

	fmt.Println("Value of A from testtime package is ", tt.A)

	math.Add(10, 20)

	math.Sub(20, 5)

	advanced.Square(100)

	//a := mavtime.A

}

func init() {

	fmt.Println("main: Init from 3rd init method")
}
