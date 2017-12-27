package main

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
)
const (
	ab = "avahi-browse"
	abargs = "-at"
)

func main() {
	res := exec.Command(ab, abargs)
	output, err := res.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	err = res.Start()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(output)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	err = res.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
