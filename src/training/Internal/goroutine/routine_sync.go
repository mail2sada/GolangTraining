package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int64

var waitGroup sync.WaitGroup
var mutex sync.Mutex

func FirstRoutine(msg string) {

	defer waitGroup.Done()

	fmt.Println("In my First Routine")

	for {
		fmt.Println("FirstRoutine: Increamneting counter:", msg)
		mutex.Lock()
		{
			counter++
		}
		mutex.Unlock()
		runtime.Gosched()
		if counter >= 30000 {
			break
		}

	}
}

func SecondRoutine(str string) {
	defer waitGroup.Done()
	fmt.Println("In My Second Routine")

	for {

		fmt.Println("SecondRoutine: Displaying counter:", str)
		mutex.Lock()
		{
			fmt.Println("Counter value is ", counter)
		}
		mutex.Unlock()
		runtime.Gosched()
		if counter >= 30000 {
			break
		}
	}
}

func main() {

	waitGroup.Add(4)

	fmt.Println("Demo: Routine Synchronization")

	go FirstRoutine("Guru")

	go SecondRoutine("Suresh")

	go SecondRoutine("Gyanendra")

	go FirstRoutine("Vishal")

	waitGroup.Wait()
}
