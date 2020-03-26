package main

import "fmt"

func main() {
	var x int = 5
	fmt.Println("Value without pointer:   ", x)
	var xptr = &x
	fmt.Println("Address pointer: ", xptr)
	fmt.Println("Access to value from pointer:", *(xptr))
}
