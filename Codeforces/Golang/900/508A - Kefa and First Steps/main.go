package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	if _, err := fmt.Fscan(reader, &n); err != nil {
		return
	}
	var maxLen, currentLen int
	var prev, current int

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &current)

		if i == 0 || current >= prev {
			currentLen++
		} else {
			currentLen = 1
		}

		if currentLen > maxLen {
			maxLen = currentLen
		}
		prev = current
	}

	fmt.Fprintln(writer, maxLen)

}
