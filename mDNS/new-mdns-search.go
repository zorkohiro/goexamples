package main

import (
	"log"
	"os"
	"time"

	"github.com/oleksandr/bonjour"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: mdns-search key [key....]")
	}

	for {
		ch := make(chan string)
		go func() {
			resolver, err := bonjour.NewResolver(nil)
			if err != nil {
				log.Fatal("Failed to initialize resolver:", err.Error())
			}

			results := make(chan *bonjour.ServiceEntry)
			go func(results chan *bonjour.ServiceEntry, exitCh chan<- bool) {
				for e := range results {
					ch <- e.Instance
				}
				exitCh <- true
			}(results, resolver.Exit)

			err = resolver.Browse("_staff._tcp", "", results)
			if err != nil {
				log.Println("Failed to browse:", err.Error())
			}
			select {}
			close(ch)
		}()
		log.Print(<- ch)
		time.Sleep(5 * time.Second)
	}
}
