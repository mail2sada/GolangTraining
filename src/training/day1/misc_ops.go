package main

/*
&: This operator returns the address of the variable.

*: This operator provides pointer to a variable.

<-:The name of this operator is receive. It is used to receive a value from the channel.
*/

import "fmt"

func main() {
	a := 4

	// Using address of operator(&) and
	// pointer indirection(*) operator
	b := &a
	fmt.Println(*b)
	*b = 7
	fmt.Println(a)
}
