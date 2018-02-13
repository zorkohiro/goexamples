package main

import (
	"fmt"
	"github.com/hashicorp/mdns"
	"time"
)

func main() {

	for {
		ech := make(chan *mdns.ServiceEntry)

		go func() {
			for entry := range ech {
				fmt.Printf("Got new entry: %v\n", entry)
			}
		}()

		err := mdns.Lookup("_workstation._tcp", ech)
		if err != nil {
			fmt.Println(err)
		}
		close(ech)
		ech = nil
		fmt.Println("----------------------------------------")
		time.Sleep(2 * time.Second)
	}
}
