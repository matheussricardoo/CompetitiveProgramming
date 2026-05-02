package main

import (
	"fmt"
)

func main() {
	var k, n, w int
	fmt.Scan(&k, &n, &w)
	solutions := k * (w * (w + 1) / 2)
	borrowed := solutions - n

	if borrowed < 0 {
		fmt.Println(0)
	} else {
		fmt.Println(borrowed)
	}
}
