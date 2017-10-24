package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f := os.Stdin
	o := os.Stdout

	h := sha256.New()
	b := make([]byte, 1<<20)

	for {
		n, err := f.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		o.Write(b[:n])
		h.Write(b[:n])
	}
	fmt.Fprintf(os.Stderr, "%x\n", h.Sum(nil))
}
