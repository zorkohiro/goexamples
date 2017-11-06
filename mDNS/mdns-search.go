package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/richtr/mdns"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: mdns-search key [key....]")
	}
	for {
		ech := make(chan *mdns.ServiceEntry)
		qp := mdns.DefaultParams(os.Args[1])
		qp.Entries = ech
		qp.Timeout = 30 * time.Second
		go func() {
			err := mdns.Query(qp)
			if err != nil {
				log.Fatal(err)
			}
			close(ech)
		}()
		for {
			en := <-ech
			if en == nil {
				break
			}
			fmt.Println(en.Host, en.AddrV4, en.AddrV6, en.Port)
		}
		log.Println("-------------------------------------")
	}

}
