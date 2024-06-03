package main

import "fmt"

func main() {

	basics()

}

func basics() {
	var arr [3]int        // array of 3 ints, all values in array initialized to 0
	arr = [3]int{1, 2, 3} // assigning values
	fmt.Println(arr[1])   // 2
	arr[1] = 99
	fmt.Println(arr)      // 1, 99, 3
	fmt.Println(len(arr)) // 3

	arr1 := [3]string{"foo", "bar"}
	arr2 := arr1
	fmt.Println(arr2) // foo bar

	arr1[0] = "hi"
	fmt.Println(arr2) // still foo bar, because assigning arrays in go performs a copy operation. arr2 is a copy of the original array.

	fmt.Println(arr1 == arr2) // false
}
