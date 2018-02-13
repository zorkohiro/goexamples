package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Alpha is a Beta Dog"
	t := "Alpha is a Beta Dog today"
	if len(s) == len(t) {
		if strings.Compare(s, t) == 0 {
			fmt.Println("strings equal")
		}
	} else if len(s) < len(t) {
		if strings.Compare(s, t[:len(s)]) == 0 {
			fmt.Println("strings equal with t longer")
		}
	} else {
		if strings.Compare(s[:len(t)], t) == 0 {
			fmt.Println("strings equal with s longer")
		}
	}
}
