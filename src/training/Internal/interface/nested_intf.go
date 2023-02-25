package main

import "fmt"

type int_base interface {
	f1()
	f2()
}

type int_derived interface {
	int_base
	f3()
}

type struct1 struct {
}

type struct2 struct {
}

func (s struct1) f1() {

}

func (s struct1) f2() {

}

func (s struct1) f3() {

}

func (s struct2) f1() {

}

func (s struct2) f2() {

}

func main() {

	fmt.Println("Demo: nested interface")

	var b int_base

	var a int_derived

	a = struct1{}

	b = struct2{}

	b = a

}
