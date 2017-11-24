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
	log.Println("youpc", addrs)
}
