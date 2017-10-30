package main

import (
	"log"
	"net"
)


func main() {
	alist, err := net.InterfaceAddrs(); if err != nil {
		log.Fatal(err)
	}

	for _, addr := range alist {
		log.Println(addr.Network(), addr.String())
	}
}
