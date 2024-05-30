package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Waddup")
	// os.Stdin forces us to read inputs one by one, so instead we use bufio
	in := bufio.NewReader(os.Stdin) // bufio is a decorator wrapping around os.Stdin without significantly changing its interface
	s, _ := in.ReadString('\n')     // new line is the delimiter. Function returns a string and an error
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)
	fmt.Println(s)
}
