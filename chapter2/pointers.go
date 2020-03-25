package main

import "fmt"

func main() {
	var x int = 5
	fmt.Println("Value:   ", x)
	var xptr = &x
	fmt.Println("Address: ", xptr)
}
