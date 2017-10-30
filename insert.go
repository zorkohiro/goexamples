package main

import (
	"fmt"
	"os"
)

func main() {
	var strs []string
	var i, j, k int

	max := len(os.Args)
	for i = 1; i < max; i++ {
		makeone := true
		slen := len(strs)
		for j = 0; j < slen; j++ {
			if os.Args[i] < strs[j] {
				// fmt.Println("insert", os.Args[i], "before", strs[j])
				l := len(strs) + 1
				nstrs := make([]string, l)
				for k = 0; k < j; k++ {
					nstrs[k] = strs[k]
				}
				nstrs[k] = os.Args[i]
				for k = k + 1; k < l; k++ {
					nstrs[k] = strs[k-1]
				}
				strs = nstrs
				makeone = false
				break
			}
		}
		if makeone {
			strs = append(strs, os.Args[i])
			// fmt.Println("append", os.Args[i], "to end")
		}
	}
	for i := 0; i < len(strs); i++ {
		fmt.Println(strs[i])
	}
}
