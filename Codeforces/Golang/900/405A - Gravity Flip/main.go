package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	if !scanner.Scan() {
		return
	}
	n, _ := strconv.Atoi(scanner.Text())

	columns := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		columns[i], _ = strconv.Atoi(scanner.Text())
	}

	slices.Sort(columns)

	for i, val := range columns {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(val)
	}
	fmt.Println()
}
