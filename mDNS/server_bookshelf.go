package main

import (
	"github.com/hashicorp/mdns"
	"log"
	"net"
	"os"
	"strings"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: mdns-advertise key")
	}

	l, err := net.Listen("tcp", ":")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	_, pstring, err := net.SplitHostPort(l.Addr().String())
	if err != nil {
		log.Fatal(err)
	}
	port, err := strconv.Atoi(pstring)
	if err != nil {
		log.Fatal(err)
	}

	host, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	info := []string{"My awesome service"}
	service, err := mdns.NewMDNSService(host, os.Args[1], "", "", port, nil, info)
	if err != nil {
		log.Fatal(err)
	}

	server, err := mdns.NewServer(&mdns.Config{Zone: service, LogEmptyResponses: false})
	if err != nil {
		log.Fatal(err)
	}
	defer server.Shutdown()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(conn)
		conn.Close()
	}
}
