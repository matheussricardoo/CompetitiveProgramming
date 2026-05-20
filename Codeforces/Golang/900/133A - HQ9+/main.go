package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	if !scanner.Scan() {
		return
	}

	s := scanner.Text()

	for i := 0; i < len(s); i++ {
		if s[i] == 'H' || s[i] == 'Q' || s[i] == '9' {
			fmt.Println("YES")
			return
		}
	}
	fmt.Println("NO")
}
