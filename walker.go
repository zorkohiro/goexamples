package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: ", os.Args[0], "directory")
	}
	rootdir := os.Args[1]

	err := filepath.Walk(rootdir, walker)
	if err != nil {
		log.Fatal(err)
	}
}

func walker(path string, info os.FileInfo, err error) error {
	if err != nil {
		return filepath.SkipDir
	}
	if info.IsDir() {
		return nil
	}
	s := fmt.Sprintf("size %10v mode: %v modtime: %v", info.Size(), info.Mode(), info.ModTime().UnixNano())
	log.Printf("% 40s: %s\n", s, path)
	return nil
}
