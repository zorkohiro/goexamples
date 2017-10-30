package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var list []string

	for i := 0; i < 5; i++ {
		list = append(list, "valuevoter" + strconv.Itoa(i))
	}
	fmt.Println("list", list)
	fred := list
	fmt.Println("fred", fred)
	john := list[0]
	fmt.Println("john", john, reflect.TypeOf(john))
	zb := list[2:4]
	fmt.Println("zb", zb, reflect.TypeOf(zb))
}
