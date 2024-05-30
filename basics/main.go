// go modules help manage dependencies in Go projects. Specifies dependencies and required versions. (go mod init)
// gopath is env variable pointing to location of go workspace containing all neccessary files for Go development.
// goroot points to the location of go installation directory (SDK).
package main // package identifier indicates that we are creating a source file for something we are going to compile later

import "fmt" // fmt package allows us to work with strings

func main() { // entry point in go

	fmt.Println("Hello Gophers!")
}
