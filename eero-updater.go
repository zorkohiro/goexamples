package main

import (
	"archive/tar"
	"crypto/sha256"
	"io"
	"log"
	"os"
)

func main() {
	var odir string

	switch len(os.Args) {
	case 1:
		odir = "."
	case 2:
		odir = os.Args[1]
	default:
		log.Fatal("usage: eero-updater [tmp-dir]")
	}

	//
	// We read from standard input
	//
	streamin := os.Stdin

	//
	// We write down a pipe to a function that extracts a tarball
	//
	pipe_in, pipe_out := io.Pipe()

	go func() {
		tr := tar.NewReader(pipe_in)
		for {
			hdr, err := tr.Next()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			switch hdr.Typeflag {
			case tar.TypeDir:
			case tar.TypeReg:
				log.Println("create", hdr.Name)
				fil, err := os.Create(odir + hdr.Name)
				if err != nil {
					log.Fatal(err)
				}
				_, err = io.Copy(fil, tr)
				if err != nil {
					log.Fatal(err)
				}
				fil.Close()
			default:
				log.Fatal("unable to figure out filetype", hdr)
			}
		}
		pipe_in.Close()
	}()

	//
	// We create a sha256 instance for checking our incoming data
	//
	sha := sha256.New()

	//
	// We have a small buffer we stage so that sha256 can look at
	// it before we unzip it and send it down to the tar function
	//
        b := make([]byte, 4096)

        for {
                n, err := streamin.Read(b)
                if err != nil {
                        if err == io.EOF {
				log.Println("eof on input")
				pipe_out.Close()
                                break
                        }
                        log.Fatal(err)
                }
                sha.Write(b[:n])
		pipe_out.Write(b[:n])
        }
        log.Printf("%x\n", sha.Sum(nil))
}
