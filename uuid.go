package main

import (
	"github.com/satori/go.uuid"
	"log"
)

func main() {
	for i := 0; i < 40; i++ {
		log.Println(uuid.NewV1())
	}
}
