package main

import "fmt"

func main() {
	// airthmatic operators
	// +, -, *, / and %

	// Addition: The ‘+’ operator adds two operands. For example, x+y.
	// Subtraction: The ‘-‘ operator subtracts two operands. For example, x-y.
	// Multiplication: The ‘*’ operator multiplies two operands. For example, x*y.
	// Division: The ‘/’ operator divides the first operand by the second. For example, x/y.
	// Modulus: The ‘%’ operator returns the remainder when the first operand is divided by the second. For example, x%y.

	/*
		Note: -, +, !, &, *, <-, and ^ are also known as unary operators and the precedence of unary operators is higher.
		++ and — operators are from statements they are not expressions, so they are out from the operator hierarchy.
	*/

	var a, b int = 20, 10

	c := a + b
	fmt.Println("Value of (a+b) is ", c)

	d := a - b
	fmt.Println("Value of (a-b) is ", d)

	e := a * b
	fmt.Println("Value of (a*b) is ", e)

	f := a / b
	fmt.Println("Value of (a/b) is ", f)

	g := a % b
	fmt.Println("Value of (a%b) is ", g)

	// relational operators
	// <, >, ==, !=, <=, >=

	/*
		‘=='(Equal To) operator checks whether the two given operands are equal or not. If so, it returns true. Otherwise, it returns false. For example,
		5==5 will return true.
		‘!='(Not Equal To) operator checks whether the two given operands are equal or not. If not, it returns true. Otherwise, it returns false.
		It is the exact boolean complement of the ‘==’ operator. For example, 5!=5 will return false.
		‘>'(Greater Than)operator checks whether the first operand is greater than the second operand. If so, it returns true. Otherwise,
		it returns false. For example, 6>5 will return true.
		‘<‘(Less Than)operator checks whether the first operand is lesser than the second operand. If so, it returns true. Otherwise, it returns false.
		For example, 6<5 will return false.
		‘>='(Greater Than Equal To)operator checks whether the first operand is greater than or equal to the second operand. If so, it returns true.
		Otherwise, it returns false. For example, 5>=5 will return true.
		‘<='(Less Than Equal To)operator checks whether the first operand is lesser than or equal to the second operand. If so, it returns true. Otherwise,
		it returns false. For example, 5<=5 will also return true.
	*/

	res := (a < b)
	fmt.Println(res)

	res1 := (a > b)
	fmt.Println(res1)

	res2 := (a == b)
	fmt.Println(res2)

	res3 := (a != b)
	fmt.Println(res3)

	res4 := (a <= b)
	fmt.Println(res4)

	res5 := (a >= b)
	fmt.Println(res5)

	//Logical operators
	// Logical AND: The ‘&&’ operator returns true when both the conditions in consideration are satisfied. Otherwise it returns false. For example, a && b returns true when both a and b are true (i.e. non-zero).
	// Logical OR: The ‘||’ operator returns true when one (or both) of the conditions in consideration is satisfied. Otherwise it returns false. For example, a || b returns true if one of a or b is true (i.e. non-zero). Of course, it returns true when both a and b are true.
	// Logical NOT: The ‘!’ operator returns true the condition in consideration is not satisfied. Otherwise it returns false. For example, !a returns true if a is false, i.e. when a=0.

	res6 := (res && res1)
	fmt.Println(res6)

	res7 := (res2 || res3)
	fmt.Println(res7)

	res8 := (!res4)
	fmt.Println(res8)

	//bitwiae operators

	// 	& (bitwise AND): Takes two numbers as operands and does AND on every bit of two numbers. The result of AND is 1 only if both bits are 1.
	// | (bitwise OR): Takes two numbers as operands and does OR on every bit of two numbers. The result of OR is 1 any of the two bits is 1.
	// ^ (bitwise XOR): Takes two numbers as operands and does XOR on every bit of two numbers. The result of XOR is 1 if the two bits are different.
	// << (left shift): Takes two numbers, left shifts the bits of the first operand, the second operand decides the number of places to shift.
	// >> (right shift): Takes two numbers, right shifts the bits of the first operand, the second operand decides the number of places to shift.
	// &^ (AND NOT): This is a bit clear operator.

}
