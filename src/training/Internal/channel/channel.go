package main

import "fmt"

func MyRoutine(pi chan int) {

	vl := 0
	vl = <-pi

	fmt.Println("Data received on chan is ", vl)

	vl++

	pi <- vl

}

func main() {

	var pipe chan int = make(chan int)

	var val int = 0

	go MyRoutine(pipe)

	pipe <- val // is a blocking in case of unbuffered channels

	val = <-pipe

	fmt.Println("Value received from Myroutine is ", val)

}
