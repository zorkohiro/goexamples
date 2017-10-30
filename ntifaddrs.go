// print a list of our non-trivial interface addresses
package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("unable to get a list of our interface addresses")
		return
	}

	for _, iaddr := range addrs {
		// remove any trailing netmask
		sl := strings.Split(iaddr.String(), "/")
		addr := sl[0]
		switch addr {
		case "127.0.0.1":
		case "::1":
		default:
			if strings.HasPrefix(addr, "fe80::") == false {
				fmt.Println(addr)
			}
		}
	}
}
