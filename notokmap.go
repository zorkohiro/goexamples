package main

import (
	"fmt"
	"reflect"
)

type zop struct {
	i	int
}

func main() {
	var n string

	m := make(map[string]string)
	m["fred"] = "bob"
	z, ok := m["fred"]
	if !ok {
		fmt.Println("couldn't find bob")
		return
	}
	fmt.Println("found fred", z)
	n, ok = m["zippy"]
	if ok {
		fmt.Println("shouldn't have been able to find zippy")
		return
	}
	fmt.Println("didn't find zippy, n is:<", n, ">", reflect.TypeOf(n))
	zz := make(map[string]*zop)
	zz["xbar"] = &zop{i: 23}
	x, ok := zz["xbar"]
	if !ok {
		fmt.Println("couldn't find xbar, x is <", x, "<", reflect.TypeOf(x))
		return
	}
	fmt.Println("found xbar, x is <", x, "<", reflect.TypeOf(x))
	y, ok := zz["nonsense"]
	if ok {
		fmt.Println("shouldn't have found nonsense, y is <", y, "<", reflect.TypeOf(y))
		return
	}
	fmt.Println("didn't find nonsense, y is:<", y, ">", reflect.TypeOf(y))
	k, ok := zz[""]
	if ok {
		fmt.Println("shouldn't have found ankthing, k is <", k, "<", reflect.TypeOf(k))
		return
	}
	fmt.Println("found nothing, k is:<", k, ">", reflect.TypeOf(k))
}
