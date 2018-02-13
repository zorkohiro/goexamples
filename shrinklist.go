package main

import (
	"fmt"
)

func main() {
	var input []string
	for i := 0; i < 10; i++ {
		s := fmt.Sprintf("string%d", i)
		input = append(input, s)
	}
	fmt.Println(input)
	input = append(input[:0], input[1:]...)
	fmt.Println(input)
}
