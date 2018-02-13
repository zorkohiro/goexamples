package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/peer"
	"log"
	"net"
)

func main() {
	var ctx context.Context
	giggle, err := net.Dial("tcp", "8.8.8.8:53")
	if err != nil {
		log.Fatal(err)
	}
	ctx = peer.NewContext(ctx, &peer.Peer{Addr: giggle.RemoteAddr()})
	log.Println(ctx)
}
