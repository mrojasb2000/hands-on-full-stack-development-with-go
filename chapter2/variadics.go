package main

import "fmt"

func infiniteAdder(inputs ...int) (sum int) {
	for _, v := range inputs {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(infiniteAdder(1, 2, 2, 2))
}
