package main

import (
	"fmt"
	"os"
	"time"
)

var looper int
var fredchan chan string
var bobchan chan int

func foo() {
	for {
		select {
		case c1 := <-fredchan:
			fmt.Println("fred got", c1)
			if c1 == "exit" {
				fmt.Println("looper", looper)
				os.Exit(0)
			}
		case c2 := <-bobchan:
			fmt.Println("bob got", c2)
		}
		looper++
	}
}

func main() {
	fredchan = make(chan string, 100)
	bobchan = make(chan int, 100)
	go foo()

	fredchan <- "fred"
	time.Sleep(time.Second)
	bobchan <- 1
	time.Sleep(5 * time.Second)
	fredchan <- "exit"
	time.Sleep(time.Second)
	for {
		time.Sleep(time.Second)
		fmt.Println(looper)
	}
}
