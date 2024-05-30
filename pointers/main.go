package main

import "fmt"

func main() { // entry point in go

	// Values
	p1 := 42
	p2 := p1
	// both 42 now
	// but if i do
	p1 = 27
	fmt.Println(p2) // still gonna be 42

	// Pointers
	p3 := 42
	p4 := &p3        // p4 holds the address of p3, essentially pointing to p3 (ref)
	fmt.Println(*p4) // 42 (deref)
	p3 = 21
	fmt.Println(*p4) // 21

	*p4 = 30        // now p3's value is 30, we can dereference pointer and change value of what its pointing to
	fmt.Println(p3) // 30

	// we can also create pointers using the keyword new.
	p := new(string)
	fmt.Println(p, *p) // new address, empty string

}
