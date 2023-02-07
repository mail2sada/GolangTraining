package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	fmt.Println("Demo: channels..")

	var channel chan int = make(chan int)

	go send(channel)
	go receive(channel)

	wg.Wait()
}

func send(ch chan int) {
	defer wg.Done()
	var sendItems = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, -1, 10, 11, 16}

	for _, val := range sendItems {
		fmt.Println("Sending data on channel ", val)
		ch <- val
		if val == -1 {
			break
		}
	}

}

func receive(ch chan int) {
	defer wg.Done()

	for {
		res, _ := <-ch

		fmt.Printf("Type of res is %T and value of res is %d\n", res, res)
		if res == -1 {
			break
		}
	}

}
