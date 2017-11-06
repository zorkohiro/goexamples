package main
import (
	"time"
	"fmt"
	"strconv"
)

func threat(pipe chan<- string) {
	for i := 1; i <= 10; i++ {
		pipe <- "threat " + strconv.Itoa(i)
		time.Sleep(500 * time.Millisecond)
	}
	close(pipe)
}

func startit() (<- chan string) {
	threats := make(chan string)
	go threat(threats)
	return threats
	
}

func main() {
	t := startit()

	i := 1
	for {
		s := <-t
		if s == "" {
			break
		}
		fmt.Println(i, s)
		i++
	}
}
