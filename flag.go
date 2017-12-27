package main
import (
	"fmt"
	"flag"
)

var debug bool

func init() {
	flag.BoolVar(&debug, "debug", false, "run in debug mode")
}

func main() {
	flag.Parse()
	fmt.Println("debug is", debug)
	fmt.Println("rest of args is:", flag.Args())
}
