package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f := os.Stdin

	gzf, err := gzip.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}
	tr := tar.NewReader(gzf)

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		name := hdr.Name

		switch hdr.Typeflag {
		case tar.TypeDir:
			fmt.Println("Directory    Name:", name)
			continue
		case tar.TypeReg:
			fmt.Println("Regular File Name:", name)
		default:
			log.Fatal("unable to figure out filetype", hdr)
		}
	}
}
