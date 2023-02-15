package main

import "fmt"

func myRoutine(a chan int) {

	fmt.Println("sending value 10")
	a <- 10
	fmt.Println("sent value 10")

}

func main() {

	fmt.Println("Demo: Channels")

	var channel = make(chan int)

	go myRoutine(channel)
	fmt.Println("receiving from routine")
	fromRoutine := <-channel
	fmt.Println("received  from routine")
	fmt.Println(fromRoutine)

}
