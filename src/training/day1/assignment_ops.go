package main

/*
“=”(Simple Assignment): This is the simplest assignment operator. This operator is used to assign the value on the right to the variable on the left.

“+=”(Add Assignment): This operator is a combination of ‘+’ and ‘=’ operators. This operator first adds the current value of the variable on left to
the value on the right and then assigns the result to the variable on the left.

“-=”(Subtract Assignment): This operator is a combination of ‘-‘ and ‘=’ operators. This operator first subtracts the current value of the variable on
left from the value on the right and then assigns the result to the variable on the left.

“*=”(Multiply Assignment): This operator is a combination of ‘*’ and ‘=’ operators. This operator first multiplies the current value of the variable on
left to the value on the right and then assigns the result to the variable on the left.

“/=”(Division Assignment): This operator is a combination of ‘/’ and ‘=’ operators. This operator first divides the current value of the variable on left
by the value on the right and then assigns the result to the variable on the left.

“%=”(Modulus Assignment): This operator is a combination of ‘%’ and ‘=’ operators. This operator first modulo the current value of the variable on left by
the value on the right and then assigns the result to the variable on the left.

“&=”(Bitwise AND Assignment): This operator is a combination of ‘&’ and ‘=’ operators. This operator first “Bitwise AND” the current value of the variable
on the left by the value on the right and then assigns the result to the variable on the left.

“^=”(Bitwise Exclusive OR): This operator is a combination of ‘^’ and ‘=’ operators. This operator first “Bitwise Exclusive OR” the current value of the
variable on left by the value on the right and then assigns the result to the variable on the left.

“|=”(Bitwise Inclusive OR): This operator is a combination of ‘|’ and ‘=’ operators. This operator first “Bitwise Inclusive OR” the current value of the
variable on left by the value on the right and then assigns the result to the variable on the left.

“<<=”(Left shift AND assignment operator): This operator is a combination of ‘<<’ and ‘=’ operators. This operator first “Left shift AND” the current value
of the variable on left by the value on the right and then assigns the result to the variable on the left.

“>>=”(Right shift AND assignment operator): This operator is a combination of ‘>>’ and ‘=’ operators. This operator first “Right shift AND” the current
value of the variable on left by the value on the right and then assigns the result to the variable on the left.
*/

import "fmt"

func main() {
	var p int = 45
	var q int = 50

	// “=”(Simple Assignment)
	p = q
	fmt.Println(p)

	// “+=”(Add Assignment)
	p += q
	fmt.Println(p)

	//“-=”(Subtract Assignment)
	p -= q
	fmt.Println(p)

	// “*=”(Multiply Assignment)
	p *= q
	fmt.Println(p)

	// “/=”(Division Assignment)
	p /= q
	fmt.Println(p)

	// “%=”(Modulus Assignment)
	p %= q
	fmt.Println(p)

}
