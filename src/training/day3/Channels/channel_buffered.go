package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

var wg sync.WaitGroup

type PlayerError struct {
	error_code  int
	description string
}

func main() {
	wg.Add(2)
	fmt.Println("Demo: Buffered channels")

	var command chan string = make(chan string, 10)

	var status chan PlayerError = make(chan PlayerError)
	go controller(command, status)
	go player(command, status)
	fmt.Println(command)
	wg.Wait()
}

func controller(command chan string, status chan PlayerError) {
	defer wg.Done()

	var cmd = [...]string{`play`, `pause`, `stop`, `start`, `exit`}

	for _, val := range cmd {
		fmt.Println("sending command :", val)
		command <- val
		if reflect.DeepEqual(command, `exit`) {
			break
		}
	}

}

func player(command chan string, status chan PlayerError) {
	defer wg.Done()
	time.Sleep(2 * time.Second)

	for {

		val := <-command

		fmt.Println("received command...", val)
		if reflect.DeepEqual(val, `exit`) {
			break
		}

	}

}
