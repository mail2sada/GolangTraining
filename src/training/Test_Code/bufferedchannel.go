package main

import (
	"fmt"
	"time"
)

func MyRoutine(ch chan int) {

	fmt.Println("In my Go routine...")

	time.Sleep(10 * time.Second)

	for {

		val := <-ch
		fmt.Println("this is from channel", val)

	}
}

func main() {

	fmt.Println("Demo: Buffere channels..")
	channel := make(chan int, 10)

	go MyRoutine(channel)

	channel <- 1
	fmt.Println("1")
	channel <- 2
	fmt.Println("2")
	channel <- 3
	fmt.Println("3")
	channel <- 4
	fmt.Println("4")
	channel <- 5
	fmt.Println("5")
	channel <- 6
	fmt.Println("6")
	channel <- 7
	fmt.Println("7")
	channel <- 8
	fmt.Println("8")
	channel <- 9
	fmt.Println("9")
	channel <- 10
	fmt.Println("10")
	channel <- 11
	fmt.Println(11)

}
