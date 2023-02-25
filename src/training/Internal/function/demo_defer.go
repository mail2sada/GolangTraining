package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("test.txt")
	if err != nil {
		return
	}

	defer func() {
		file.Close()
	}()

	fmt.Fprintln(file, "We are wrtining in to this file.....")

}
