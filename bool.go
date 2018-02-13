package main

import (
	"fmt"
)

func bob(fred bool) bool {
	return !fred
}
func main() {
	fmt.Println(bob(false))
	fmt.Println(bob(true))
}
