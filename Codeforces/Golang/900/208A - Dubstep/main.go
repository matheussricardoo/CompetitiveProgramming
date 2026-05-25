package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	input := scanner.Text()
	replaced := strings.ReplaceAll(input, "WUB", " ")
	words := strings.Fields(replaced)
	fmt.Println(strings.Join(words, " "))
}
