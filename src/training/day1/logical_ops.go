package main

/*
Logical AND: The ‘&&’ operator returns true when both the conditions in consideration are satisfied. Otherwise it returns false.
For example, a && b returns true when both a and b are true (i.e. non-zero).

Logical OR: The ‘||’ operator returns true when one (or both) of the conditions in consideration is satisfied. Otherwise it returns
false. For example, a || b returns true if one of a or b is true (i.e. non-zero). Of course, it returns true when both a and b are true.

Logical NOT: The ‘!’ operator returns true the condition in consideration is not satisfied. Otherwise it returns false. For example,
!a returns true if a is false, i.e. when a=0.
*/

import "fmt"

func main() {
	var p int = 23
	var q int = 60

	if p != q && p <= q {
		fmt.Println("True")
	}

	if p != q || p <= q {
		fmt.Println("True")
	}

	if !(p == q) {
		fmt.Println("True")
	}

}
