package main
import (
	"fmt"
	"time"
)
func main() {
	t := time.Now()
	fmt.Printf("now is: %v\n", t)
	fmt.Printf("now is: %v\n", t.Local())
	fmt.Printf("now is: %v\n", t.UnixNano())
	//fmt.Printf("now is: %s\n", t.Format("Mon Jan 2 15:04:05 -0000 UTC 2006"))
	fmt.Printf("now is: %s\n", t.Format("Mon Jan 2 15:04:05 2006"))
	i := 7
	fmt.Println("sleeping", i, "seconds on", time.Now().Format("Mon Jan 2 15:04:05 2006"))
	time.Sleep(time.Duration(i) * time.Second)
	fmt.Println("        sleep done on", time.Now().Format("Mon Jan 2 15:04:05 2006"))
}
