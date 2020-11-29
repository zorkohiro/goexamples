package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: ", os.Args[0], "path")
	}
	rootdir := os.Args[1]

	path, err := filepath.EvalSymlinks(rootdir)
	if err != nil {
		log.Fatal(err)
	} else {
		print(path, "\n")
	}
}
