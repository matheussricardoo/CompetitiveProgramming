package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(n int) int {
	c2 := 0
	c3 := 0

	for n%2 == 0 {
		c2++
		n /= 2
	}

	for n%3 == 0 {
		c3++
		n /= 3
	}

	if n != 1 || c2 > c3 {
		return -1
	}

	return (c3 - c2) + c3
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	if _, err := fmt.Fscan(reader, &t); err != nil {
		return
	}

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(reader, &n)
		fmt.Fprintln(writer, solve(n))
	}
}
