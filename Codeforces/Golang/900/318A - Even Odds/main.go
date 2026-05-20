package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	if !scanner.Scan() {
		return
	}
	n, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	if !scanner.Scan() {
		return
	}
	k, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	var odd int64
	if n%2 == 0 {
		odd = n / 2
	} else {
		odd = (n + 1) / 2
	}

	if k <= odd {
		fmt.Println(2*k - 1)
	} else {
		k = k - odd
		fmt.Println(2 * k)
	}
}
