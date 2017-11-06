package main

import (
	"log"
	"net"
	"strconv"
	"time"

	"github.com/oleksandr/bonjour"
)

const (
	service	= "_spasebow._tcp"
)

func main() {
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

	s, err := bonjour.Register("Spasebow", service, "local", port, []string{"txtv=1", "app=test"}, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Shutdown()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(conn)
		conn.Close()
		time.Sleep(100 * time.Millisecond)
	}
}
