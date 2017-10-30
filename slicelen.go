package main

import (
	"fmt"
)

func main() {
	var fred []int

	fmt.Println("len is", len(fred))
	fred = append(fred, 23)
	fmt.Println("len is", len(fred))
	for i := 0; i < 1000; i = i + 30 {
		fred = append(fred, i)
	}
	fmt.Println("len is", len(fred))
	for i := 0; i < len(fred); i++ {
		fmt.Println(fred[i])
	}
}
