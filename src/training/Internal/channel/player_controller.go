package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var waitGroup sync.WaitGroup

func Player(command chan string) {

	defer waitGroup.Done()

	for {
		cmd := <-command

		switch cmd {
		case "init":
			fmt.Println("Initializing player...")
		case "start":
			fmt.Println("Starting player")
		case "play":
			fmt.Println("Performing play")
		case "pause":
			fmt.Println("Performing pause")
		case "stop":
			fmt.Println("Stopping player")
		case "exit":
			fmt.Println("Exiting player")
			return
		}

	}

}

func Controller(commandIntf chan string) {
	defer waitGroup.Done()

	var commands []string = []string{"init", "start", "play", "pause", "stop", "exit"}

	for {
		idx := rand.Int63n(int64(len(commands)))

		commandIntf <- commands[idx]
		if commands[idx] == "exit" {
			break
		}
	}

}

func main() {
	waitGroup.Add(2)
	var pipe chan string = make(chan string, 5)
	fmt.Println("Demo: player")
	go Player(pipe)
	go Controller(pipe)
	waitGroup.Wait()
}
