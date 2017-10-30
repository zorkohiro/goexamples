package main

import (
	"fmt"
	"net"
	"strconv"
)

func main() {
	var port int

	for port = 50000; port < 65536; port++ {
		_, err := net.Listen("tcp", ":" + strconv.Itoa(port)); if err != nil {
			fmt.Println("Barf", err)
		} else {
			break
		}
	}
	fmt.Println(port)
}
