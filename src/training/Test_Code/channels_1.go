package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func StartController(pipe chan interface{}) {
	defer wg.Done()

	commands := []string{"Start", "Play", "Pause", "Stop", "Exit"}

	pipe <- commands[0]
	response := <-pipe
	idx := 1
	for response == 1 {
		time.Sleep(1 * time.Second)
		pipe <- commands[idx]

		response = <-pipe
		switch response.(type) {
		case int:
			if response == 1 {
				fmt.Println("Command ", commands[idx], "Executed...")
			}

		case string:
			pipe <- "Exit"
		}

		idx += 1
		if idx == len(commands)-1 {
			idx = 1
		}

	}
}

func StartExecuter(pipe chan interface{}) {

	defer wg.Done()

	for {
		command := <-pipe

		switch command {

		case "Start":
			fmt.Println("Execution started..")
			pipe <- 1
		case "Stop":
			fmt.Println("Execution stopped..")
			pipe <- 1
		case "Play":
			fmt.Println("Execution play..")
			pipe <- 1
		case "Pause":
			fmt.Println("Execution pause..")
			pipe <- 1
		case "Exit":
			fmt.Println("Execution pause..")
			pipe <- 1
			return

		}

	}

}

func main() {

	wg.Add(2)

	fmt.Println("Demo: Channels...")

	var channel = make(chan interface{})

	go StartController(channel)

	go StartExecuter(channel)

	time.AfterFunc(10*time.Second, func() {
		channel <- "Exit"
	})

	wg.Wait()

}
