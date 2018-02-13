package main

import (
	"fmt"
	"time"
)

func threat(pipe chan<- string, stop <-chan bool) {
	for {
		pipe <- "threat"
		time.Sleep(1 * time.Second)
		select {
		case <-stop:
			fmt.Println("thread stop")
			close(pipe)
			return
		default:
		}
	}
}

func startit(stopit <-chan bool) <-chan string {
	threats := make(chan string)
	go threat(threats, stopit)
	return threats

}

func main() {
	stopit := make(chan bool)
	t := startit(stopit)

	for i := 0; ; i++ {
		s, ok := <-t
		if !ok {
			break
		}
		if s == "" {
			break
		}
		fmt.Println(i+1, s)
		if i+1 == 10 {
			fmt.Println("tell threat to stop")
			stopit <- true
		}
	}
}
