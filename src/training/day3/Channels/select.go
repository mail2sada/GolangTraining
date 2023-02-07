package main

import (
	"fmt"
	"time"
)

func chanOneRoutine(ch chan int) {

	fmt.Println("In chanOneRoutine...")
	for {
		time.Sleep(2 * time.Second)
		ch <- 1
	}
}

func chanTwoRoutine(ch chan int) {

	fmt.Println("In chanTwoRoutine...")
	for {
		time.Sleep(1 * time.Second)
		ch <- 2
	}
}

func main() {

	fmt.Println("Demo: Usage of select...")

	chanOne, chanTwo := make(chan int), make(chan int)

	fmt.Println(chanOne, chanTwo)

	go chanOneRoutine(chanOne)

	go chanTwoRoutine(chanTwo)

	for {
		select {
		case b := <-chanOne:
			fmt.Println("Value received from chanOne is ", b)
			break
		case a := <-chanTwo:
			fmt.Println("Value received from chanOne is ", a)
			break
		case <-time.After(3000 * time.Millisecond):
			fmt.Println("Timed out...")
		}
	}

}
