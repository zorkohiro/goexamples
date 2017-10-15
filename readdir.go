package main

import (
	"fmt"
	"os"
)

func main() {
	var ndir, j int

	fp, err := os.Open(os.Args[1]); if err != nil {
		fmt.Println(err)
		return
	}
	fpinfo, err := fp.Readdir(0); if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < len(fpinfo); i++ {
		if fpinfo[i].IsDir() {
			ndir++
		}
	}

	dirs := make([]string, ndir)
	j = 0
	for i := 0; i < len(fpinfo); i++ {
		if fpinfo[i].IsDir() {
			dirs[j] = fpinfo[i].Name()
			j++
		}
	}
	fmt.Println(len(fpinfo), "entries", ndir, "directories")
	fmt.Println("dirs", dirs)
}
