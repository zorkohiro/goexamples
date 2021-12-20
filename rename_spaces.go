/*
 * Walk a tree and rename normal files with spaces in the filenames replaced with underscores
 */
package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
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
	if (strings.Contains(path, " ")) {
		npath := strings.Replace(path, " ", "_", -1)
		println(path, " -> ", npath)
		err := os.Rename(path, npath); if err != nil {
			log.Fatal("cannot rename: ", err)
		}
	}
	return nil
}
