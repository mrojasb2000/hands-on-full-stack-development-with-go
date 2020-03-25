package main

import "fmt"

func main() {
	var s string
	fmt.Println(s)
	var s1, s2, s3 string
	fmt.Println(s1, s2, s3)
	var s11, s22, s33 string = "first-string", "second-string", "thrid-string"
	fmt.Println(s11, s22, s33)
	var s4, s5, s6 string = "first-string", "second-string", "thrid-string"
	fmt.Println(s4, s5, s6)
	var s44, s55, s66 string = "first-string", "second-string", "thrid-string"
	fmt.Println(s44, s55, s66)
	var (
		name = "mystring"
		number = 12
		price = 14.53
	)
	fmt.Println(name, number, price)
	
	fmt.Println("Variables in go")
}