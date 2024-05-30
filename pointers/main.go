package main

import "fmt"

func main() { // entry point in go

	// Pointers
	p1 := 42
	p2 := p1
	// both 42 now
	// but if i do
	p1 = 27
	fmt.Println(p2) // still gonna be 42

	p3 := 42
	p4 := &p3        // p4 holds the address of p3, essentially pointing to p3
	fmt.Println(*p4) // 42
	p3 = 21
	fmt.Println(*p4) // 21

	*p4 = 30        // now p3's value is 30, we can dereference pointer and change value of what its pointing to
	fmt.Println(p3) // 30

}
