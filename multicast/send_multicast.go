package main

import (
	"log"
	"net"
	"strconv"
	"time"
)

func main() {
	const (
		udpaddr = "224.0.0.1:33"
		advhdr  = "HI THERE!"
	)
	port := 36356

	addr, err := net.ResolveUDPAddr("udp", udpaddr)
	if err != nil {
		log.Fatal("cannot resolve UDP address", udpaddr, "error:", err)
	}
	c, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal("could not connect to multicast udp:", err)
	}
	xmitstring := advhdr + strconv.Itoa(port)

	for {
		c.Write([]byte(xmitstring))
		time.Sleep(5 * time.Second)
	}
}
