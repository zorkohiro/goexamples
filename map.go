package main

import (
	"fmt"
)

type fred struct {
	bob int
}

type zop struct {
	name  string
	value int
}

func main() {
	var m map[string]*fred

	m = make(map[string]*fred)

	b := new(fred)
	b.bob = 23
	m["fred"] = b
	c := new(fred)
	c.bob = 42
	m["xzb"] = c
	fmt.Println(m, m["fred"], m["fred"].bob)
	fmt.Println(m, m["xzb"], m["xzb"].bob)
	fmt.Println(m, m["zilly"])
	fmt.Println(len(m))
	gorp := make(map[string]zop)
	gorp["zz"] = zop{ "zz", 23 }
	gorp["zzx"] = zop{  "zzx", 46 }
	fmt.Println(gorp["zz"], gorp["zzx"], gorp["qrt"])
	if gorp["qrt"].name == "" {
		fmt.Println("gorp[qrt] is undefined")
	}
}
