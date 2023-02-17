package main

import (
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("[LOGS_MAIN]")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Print("Logs started")

}

func main() {

	defer func() {
		fmt.Println("Exiting...")
	}()

	log.Println("Inside main function...")

	for i := 0; i < 10; i++ {

		log.Println("Executining loop iteration ", i)

		//fmt.Println(i)

		log.Println("Printining i is complete")
	}

	log.Println("Exiting main...")

	//log.Panicln("This is Panic error")

	log.Fatalln("This is a fatal error...")

	fmt.Println("Hello....")

}
