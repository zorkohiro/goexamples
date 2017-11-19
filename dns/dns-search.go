package main

import (
	"log"
	"net"
)

func main() {
	addrs, err := net.LookupHost("youpc.local")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(addrs)
	s, srv, err := net.LookupSRV("bookshelf", "tcp", "local")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(s, srv)
}
