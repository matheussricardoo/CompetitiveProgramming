package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	if _, err := fmt.Fscan(reader, &n, &m); err != nil {
		return
	}

	puzzles := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &puzzles[i])
	}

	slices.Sort(puzzles)

	minDiff := int(^uint(0) >> 1)

	for i := 0; i <= m-n; i++ {
		diff := puzzles[i+n-1] - puzzles[i]

		if diff < minDiff {
			minDiff = diff
		}
	}

	fmt.Fprintln(writer, minDiff)
}
