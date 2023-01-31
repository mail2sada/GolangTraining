package main

/*
& (bitwise AND): Takes two numbers as operands and does AND on every bit of two numbers. The result of AND is 1 only if both bits are 1.

| (bitwise OR): Takes two numbers as operands and does OR on every bit of two numbers. The result of OR is 1 any of the two bits is 1.

^ (bitwise XOR): Takes two numbers as operands and does XOR on every bit of two numbers. The result of XOR is 1 if the two bits are different.

<< (left shift): Takes two numbers, left shifts the bits of the first operand, the second operand decides the number of places to shift.

>> (right shift): Takes two numbers, right shifts the bits of the first operand, the second operand decides the number of places to shift.

&^ (AND NOT): This is a bit clear operator.
*/

import "fmt"

func main() {
	p := 34
	q := 20

	// & (bitwise AND)
	result1 := p & q
	fmt.Printf("Result of p & q = %d", result1)

	// | (bitwise OR)
	result2 := p | q
	fmt.Printf("\nResult of p | q = %d", result2)

	// ^ (bitwise XOR)
	result3 := p ^ q
	fmt.Printf("\nResult of p ^ q = %d", result3)

	// << (left shift)
	result4 := p << 1
	fmt.Printf("\nResult of p << 1 = %d", result4)

	// >> (right shift)
	result5 := p >> 1
	fmt.Printf("\nResult of p >> 1 = %d", result5)

	// &^ (AND NOT)
	result6 := p &^ q
	fmt.Printf("\nResult of p &^ q = %d", result6)

}
