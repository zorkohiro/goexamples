package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatal("usage: httpfetch url")
	}
	url := os.Args[1]
	//
	// We read from standard input
	//
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
	b := make([]byte, 4096)
	f, err := os.Create("/tmp/zjan")
	if err != nil {
		log.Fatal(err)
	}
	for {
		n, err := res.Body.Read(b)
		if err != nil {
			if err == io.EOF {
				if n != 0 {
					_, err = f.Write(b[:n])
				}
				f.Close()
				break
			}
			log.Fatal(err)
		}
		_, err = f.Write(b[:n])
		if err != nil {
			log.Fatal(err)
		}
	}
	res.Body.Close()
}
