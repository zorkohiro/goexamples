// print a list of our non-trivial interface addresses
package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "/1/2/3/4/5/mm"
	fmt.Println("path:", path)
	dirtree := strings.Split(path, "/")
        n, dirtree := dirtree[len(dirtree)-1], dirtree[1:len(dirtree)-1]
	fmt.Println("name component:", n)
	for _, s := range dirtree {
		fmt.Println(s)
	}
}
