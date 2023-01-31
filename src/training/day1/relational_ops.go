package main

/*
‘=='(Equal To) operator checks whether the two given operands are equal or not. If so, it returns true. Otherwise, it returns false. For example,
5==5 will return true.

‘!='(Not Equal To) operator checks whether the two given operands are equal or not. If not, it returns true. Otherwise, it returns false.
It is the exact boolean complement of the ‘==’ operator. For example, 5!=5 will return false.

‘>'(Greater Than)operator checks whether the first operand is greater than the second operand. If so, it returns true. Otherwise, it returns false.
For example, 6>5 will return true.

‘<‘(Less Than)operator checks whether the first operand is lesser than the second operand. If so, it returns true. Otherwise, it returns false.
For example, 6<5 will return false.

‘>='(Greater Than Equal To)operator checks whether the first operand is greater than or equal to the second operand. If so, it returns true.
Otherwise, it returns false. For example, 5>=5 will return true.

‘<='(Less Than Equal To)operator checks whether the first operand is lesser than or equal to the second operand. If so, it returns true. Otherwise,
it returns false. For example, 5<=5 will also return true.
*/

import "fmt"

func main() {
	p := 34
	q := 20

	// ‘=='(Equal To)
	result1 := p == q
	fmt.Println(result1)

	// ‘!='(Not Equal To)
	result2 := p != q
	fmt.Println(result2)

	// ‘<‘(Less Than)
	result3 := p < q
	fmt.Println(result3)

	// ‘>'(Greater Than)
	result4 := p > q
	fmt.Println(result4)

	// ‘>='(Greater Than Equal To)
	result5 := p >= q
	fmt.Println(result5)

	// ‘<='(Less Than Equal To)
	result6 := p <= q
	fmt.Println(result6)

}
