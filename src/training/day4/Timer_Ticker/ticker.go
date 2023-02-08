package main

import (
	"fmt"
	"time"
)

func ticker_handler(data chan int, interval int) {
	duration := 0
	for val := range time.Tick(time.Duration(interval) * time.Second) {

		fmt.Println(val)
		duration += 3
		data <- duration
	}
}

func main() {

	fmt.Println("Dem0: ticker....")
	ch := make(chan int)
	go ticker_handler(ch, 3)
	for {
		select {
		case val := <-ch:
			fmt.Println("ticket running sunce [", val, "] selcones")
		}
	}
}
