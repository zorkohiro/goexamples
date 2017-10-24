package main

import (
	"log"
	"net"
)

func main() {
	const (
		udpaddr = "224.0.0.1:33"
		dgsize  = 128
	)
	addr, err := net.ResolveUDPAddr("udp", udpaddr)
	if err != nil {
		log.Fatal(err)
	}

	l, err := net.ListenMulticastUDP("udp", nil, addr)
	l.SetReadBuffer(dgsize)
	for {
		b := make([]byte, dgsize)
		n, src, err := l.ReadFrom(b)
		if err != nil {
			log.Fatal("ReadFromUDP failed:", err)
		}
		s := string(b[:n])
		x, port, err := net.SplitHostPort(src.String())
		log.Println("heard:", s, "says", x, port, err)
	}
}
