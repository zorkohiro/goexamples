package main

import (
	"encoding/hex"
	"log"
	"net"
	"os"
	"time"
)

const (
	srvAddr         = "234.172.30.2:4272"
	maxDatagramSize = 8192
)

func main() {
	var ifname string
	if len(os.Args) > 1 {
		ifname = os.Args[1]
	}
	go ping(srvAddr)
	serveMulticastUDP(srvAddr, ifname, msgHandler)
}

func ping(a string) {
	hname, err := os.Hostname()
	addr, err := net.ResolveUDPAddr("udp", a)
	if err != nil {
		log.Fatal(err)
	}
	c, err := net.DialUDP("udp", nil, addr)
	for {
		c.Write([]byte(hname))
		time.Sleep(1 * time.Second)
	}
}

func msgHandler(src *net.UDPAddr, n int, b []byte) {
	log.Println(n, "bytes read from", src)
	log.Println(hex.Dump(b[:n]))
}

func serveMulticastUDP(a string, ifname string, h func(*net.UDPAddr, int, []byte)) {
	var err error
	var ifn *net.Interface

	if len(ifname) > 0 {
		ifn, err = net.InterfaceByName(ifname)
		if err != nil {
			log.Fatal("cannot find interface", ifname)
		}
	}
	addr, err := net.ResolveUDPAddr("udp", a)
	if err != nil {
		log.Fatal(err)
	}
	l, err := net.ListenMulticastUDP("udp", ifn, addr)
	l.SetReadBuffer(maxDatagramSize)
	for {
		b := make([]byte, maxDatagramSize)
		n, src, err := l.ReadFromUDP(b)
		if err != nil {
			log.Fatal("ReadFromUDP failed:", err)
		}
		h(src, n, b)
	}
}
