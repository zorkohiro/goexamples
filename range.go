package main

import (
	"fmt"
	"strconv"
)

type fred struct {
	bob  string
	bill int
}

func main() {
	var list []*fred

	for i := 0; i < 20; i++ {
		n := new(fred)
		n.bob = "value voter" + strconv.Itoa(i-100)
		n.bill = i
		list = append(list, n)
		fmt.Println(n)
	}
	fmt.Println(list)

	for i, j := range list {
		fmt.Println("Key", i, "Value", j, "Separate", j.bob, j.bill)
	}
}
