package main

import (
	"log"
	"net"
	"strings"
)

type Identity struct {
	Id        string
	Port      int32
	Addresses []string
}

var Self Identity

func init() {
	Self.Id = "fred"
	alist, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal("cannot get a list of interface addresses", err)
	}

	for _, addr := range alist {
		s := strings.Split(addr.String(), "/")[0]
		// eliminate IPv6 link scope
		if strings.Contains(s, "fe80:") {
			continue
		}
		// eliminate localhost
		if strings.Compare(s, "127.0.0.1") == 0 {
			continue
		}
		// eliminate localhost IPv6
		if strings.Compare(s, "::1") == 0 {
			continue
		}
		// eliminate RFC 3927 link local
		if strings.Compare(s[:7], "169.254") == 0 {
			continue
		}
		Self.Addresses = append(Self.Addresses, s)
	}

}
func main() {
	log.Println(Self)
}
