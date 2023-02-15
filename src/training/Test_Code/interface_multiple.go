package main

import "fmt"

type intf1 interface {
	f1() 
	f2()

}

type intf2 interface {
	f3()
	f4()

}

type intf3 interface {
	intf1
	intf2
}

type S struct {
	int
	string
	float64
}

func (a S) f1() {

}

func (a S) f2() {

}

func (a S) f3() {

}

func (a S) f4() {

}

func (a S) F5() {

}

func main() {

	fmt.Println("Demo multiple interface....")

}
