package main

import "fmt"

func main() {
	var x int = 5
	fmt.Println("Value without pointer:   ", x)
	var xptr = &x
	fmt.Println("Address pointer: ", xptr)
	// de-referencing
	y := *xptr
	fmt.Println("Access to value from pointer:", y)

	*xptr = 4
	fmt.Println("Change the value that the pointer points to:", *xptr)

}
