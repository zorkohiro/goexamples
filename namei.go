// print a list of our non-trivial interface addresses
package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	for _, p := range os.Args[1:] {
		fmt.Print(p)
		components := strings.Split(p, "/")
		lastcomp, components := components[len(components)-1], components[:len(components)-1]
		for _, cp := range components {
			fmt.Println(" ", cp)
		}
		fmt.Println("  ", lastcomp)
	}
}
