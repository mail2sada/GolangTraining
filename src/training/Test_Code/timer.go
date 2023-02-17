package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Printf("Demo: Timer demo...")

	timer := time.NewTimer(30 * time.Second)
	//timer1 := time.NewTimer(1 * time.Second)

	timer1 := time.NewTicker(1 * time.Second)

	timer2 := time.After(10 * time.Second)

	time.AfterFunc(10*time.Second, func() {
		fmt.Println("{}{}{}{}{}{}{}Timer has expired...")
	})

	for {

		select {
		case val := <-timer.C:
			fmt.Println("Timer expired....", val)
			break
		case val := <-timer1.C:
			fmt.Println("Timer[1] expired....", val)
		case val := <-timer2:
			fmt.Println("Timer [2] has expired", val)
		}
	}

}
