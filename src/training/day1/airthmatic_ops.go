package main

import "fmt"

/*
Addition: The ‘+’ operator adds two operands. For example, x+y.
Subtraction: The ‘-‘ operator subtracts two operands. For example, x-y.
Multiplication: The ‘*’ operator multiplies two operands. For example, x*y.
Division: The ‘/’ operator divides the first operand by the second. For example, x/y.
Modulus: The ‘%’ operator returns the remainder when the first operand is divided by the second. For example, x%y
*/

/*
Note: -, +, !, &, *, <-, and ^ are also known as unary operators and the precedence of unary operators is higher. ++
and — operators are from statements they are not expressions, so they are out from the operator hierarchy.
*/

func main() {
	var p = 34
	var q = 20

	// Addition
	result1 := p + q
	fmt.Printf("Result of p + q = %d", result1)

	// Subtraction
	result2 := p - q
	fmt.Printf("\nResult of p - q = %d", result2)

	// Multiplication
	result3 := p * q
	fmt.Printf("\nResult of p * q = %d", result3)

	// Division
	result4 := p / q
	fmt.Printf("\nResult of p / q = %d", result4)

	// Modulus
	result5 := p % q
	fmt.Printf("\nResult of p %% q = %d", result5)
}
