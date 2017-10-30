package main

import (
	"fmt"
	"strconv"
)

func main() {
	var fred int
	var bob string

	fred = 235
	bob = strconv.Itoa(fred)
	fmt.Println("Integer", fred, "Bob", bob)
}
