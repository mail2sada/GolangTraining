package main

import "fmt"

type CustomError struct {
	errorCode   int
	exceptionId int
	description string
}

func (err CustomError) Error() string {

	fmt.Println("This is getting called... ")

	str := fmt.Sprintf("ErrorCode:[%d] ExceptionId [%d], Description:[%s]", err.errorCode, err.exceptionId, err.description)

	return str
}

func Divide(size int, base int, list []int) (ans []int, err error) {

	if size >= len(list) {

		err = &CustomError{errorCode: 1, exceptionId: 5, description: "Array out of bounds..."}
		return
	}
	for i := 0; i < size; i++ {

		if list[i] == 0 {
			err = &CustomError{errorCode: 2, exceptionId: 8, description: "Divide by zero error"}
			return
		}

		ans = append(ans, base/list[i])
	}
	return
}

func main() {

	fmt.Println("Demo: Custom error")

	var intSlice = []int{10, 20, 30, 40, 50, 0, -1, 16, 36}

	list, err := Divide(8, 50, intSlice)

	fmt.Println(list, err)

}
