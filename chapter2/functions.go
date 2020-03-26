package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func addSubstract(a, b int) (int, int) {
	return a + b, a - b
}

func addSubstractWithName(a, b int) (add, sub int) {
	add = a + b
	sub = a - b
	return
}

var adder = func(a, b int) int {
	return a + b
}

var subtrator = func(a, b int) int {
	return a - b
}

var addResult = adder(3, 2)
var subResult = subtrator(3, 2)

func execute(op func(int, int) int, a, b int) int {
	return op(a, b)
}

func main() {
	fmt.Println(execute(adder, 3, 2))
	fmt.Println(execute(subtrator, 3, 2))
}
