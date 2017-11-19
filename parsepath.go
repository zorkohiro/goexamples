package main
import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		instr := s.Text()
		fmt.Print(instr +  " ")
		slcs := strings.Split(instr, "/")
		for i := len(slcs) - 1; i >= 0; i-- {
			fmt.Print(slcs[i] + " ")
		}
		fmt.Println("")
	}
}
