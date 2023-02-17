package main

import (
	"fmt"
)

func Divde(a, b int) (ans int, err error) {

	if b == 0 {
		//err = errors.New("Divde by zero error...")
		err = fmt.Errorf("Divide by zero....")
		ans = -1
	} else {

		ans = a / b
		err = nil
	}
	return

}

func main() {

	fmt.Println("Demo: Errors...")

	var intSlice = []int{10, 20, 30, 40, 50, 0, -1, 16, 36}

	for idx, val := range intSlice {

		fmt.Println("Idx and val ", idx, val)
		res, err := Divde(idx, val)

		fmt.Println("result is ======>", res, err)

	}

}
