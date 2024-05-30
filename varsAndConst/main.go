package main

import "fmt"

// 4 simple data types String Number Boolean Error
// Strings: raw string ``(escape sequences dont work here) and interpreted string ""
// Numbers: int, uint, float32, float64, complex numbers 64, 128
// Booleans: true false
// Error: we can define their types.

// 3 ways to define variables
// var myName string = "" (redundant)
// var myName = "" (compile infers string)
// myName := "" (short declaration syntax, commonly used, same thing as above)

// Type conversions
// no implicit conversions
// we do type casting

func main() {

	// cannot have unused local variables, compile error
	var a string = "foo"
	fmt.Println(a)

	var b = 99
	fmt.Println(b)

	d := 3.14
	fmt.Println(d)

	e := int(d)
	fmt.Println(e)

	f, g := 10, 5
	h := f + g
	fmt.Println(h)
	fmt.Println(f + g)

	i := f == g
	fmt.Println(i)

	const j = 42 // dont have to be used
	// j can be assigned to ints or floats
	const k float32 = 3 // cant be assigned to other types (explicit const)
	// we can also create groups of constants const (d = 3, e = 4, f (takes prev consts value))

	const l = iota // initialized to 0

	const (
		m = 2 * 5     // 10
		n             // 10
		o = iota      // 2
		p             // 3
		q = 10 * iota // 40
	)

	fmt.Println(j, k, l, m, n, o, p, q)
}
