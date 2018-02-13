package main

import (
	"log"
	"os"
)

func main() {
	loghdl := log.New(os.Stdout, "logger ", log.LstdFlags|log.LUTC)
	loghdl.Println("got me")
}
