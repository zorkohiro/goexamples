package main

import (
	"log"
	"os"
	"time"

	"github.com/oleksandr/bonjour"
)

const (
	service = "_workstation._tcp"
)

func main() {
	resolver, err := bonjour.NewResolver(nil)
	if err != nil {
		log.Fatal("Failed to initialize resolver:", err.Error())
	}

	results := make(chan *bonjour.ServiceEntry)

	go func(results chan *bonjour.ServiceEntry, exitCh chan<- bool) {
		for e := range results {
			log.Println("%s", e)
			exitCh <- true
			time.Sleep(1e9)
			os.Exit(0)
		}
	}(results, resolver.Exit)

	err = resolver.Browse(service, "local.", results)
	if err != nil {
		log.Println("Failed to browse:", err.Error())
	}

	select {}
}
