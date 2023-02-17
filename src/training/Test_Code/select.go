package main

import (
	"fmt"
	"runtime"
	"time"
)

func f1(ch chan int) {

	idx := 0

	for {

		ch <- idx
		idx = idx + 1
		time.Sleep(3 * time.Second)

	}

}

func f2(ch chan int) {
	idx := 0

	for {

		ch <- idx
		idx = idx + 1
		time.Sleep(1 * time.Second)

	}
}

func main() {

	var c1 = make(chan int)
	var c2 = make(chan int)

	

	go f1(c1)

	go f2(c2)

	for {

		select {
		case val := <-c1:
			fmt.Println("Value is ", val)
		case val := <-c2:
			fmt.Println("Value received is ", val)

		}
	}

}
