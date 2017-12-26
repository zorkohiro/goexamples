package main

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
)

var debug bool

func main() {
	var err error
	var tf *os.File
	if len(os.Args) > 2 {
		log.Fatal("usage: tartv.go [file]")
	}
	if len(os.Args) == 2 {
		tf, err = os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	} else {
		tf = os.Stdin
	}
	tr := tar.NewReader(tf)
	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		switch hdr.Typeflag {
		case tar.TypeLink:
			fmt.Println("ln: ", hdr.Name)
		case tar.TypeSymlink:
			fmt.Println("symlink: ", hdr.Name)
		case tar.TypeChar:
			fmt.Println("cdev: ", hdr.Name)
		case tar.TypeBlock:
			fmt.Println("bdev: ", hdr.Name)
		case tar.TypeFifo:
			fmt.Println("fifo: ", hdr.Name)
		case tar.TypeCont:
			fmt.Println("cont: ", hdr.Name)
		case tar.TypeXHeader:
			fmt.Println("TypeXHeader: ", hdr.Name)
		case tar.TypeXGlobalHeader:
			fmt.Println("TypeXGlobalHeader: ", hdr.Name)
		case tar.TypeGNULongName:
			fmt.Println("TypeGNULongName: ", hdr.Name)
		case tar.TypeGNULongLink:
			fmt.Println("TypeGNULongLink: ", hdr.Name)
		case tar.TypeGNUSparse:
			fmt.Println("TypeGNUSparse: ", hdr.Name)
		case tar.TypeDir:
			fmt.Println("dir: ", hdr.Name)
		case tar.TypeReg:
			fmt.Println("file: ", hdr.Name)
		default:
			log.Fatal("unable to figure out filetype", hdr)
		}
	}
}
