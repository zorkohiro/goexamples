package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	fhdl, err := os.Open("read.cfg")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(fhdl)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
