package main

import "fmt"

type inputError struct {
	message      string
	missingField string
}

func (i *inputError) Error() string {
	return i.message
}

func (i *inputError) getMissingField() string {
	return i.missingField
}

func main() {
	fmt.Println("Demo: Custom error...")
	err := validate("Vishal", "Male")
	if err != nil {
		if err, ok := err.(*inputError); ok {
			fmt.Println(err)
			fmt.Printf("Missing Field is %s\n", err.getMissingField())
		}
	}
}

func validate(name, gender string) error {
	if name == "" {
		return &inputError{message: "Name is mandatory", missingField: "name"}
	}
	if gender == "" {
		return &inputError{message: "Gender is mandatory", missingField: "gender"}
	}
	return nil
}
