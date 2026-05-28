package main

import (
	"bufio"
	"fmt"
	"os"
)

func getVinner(n, m int) string {
	min := n
	if m < n {
		min = m
	}

	if min%2 != 0 {
		return "Akshat"
	}
	return "Malvika"
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	if _, err := fmt.Fscan(reader, &n, &m); err != nil {
		return
	}

	fmt.Fprintln(writer, getVinner(n, m))
}
