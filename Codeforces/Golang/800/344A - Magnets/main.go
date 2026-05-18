package main

import (
	"fmt"
)

func main() {
	var n, current, previous, count int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {

		fmt.Scan(&current)
		if current != previous {
			count++
		}
		previous = current
	}
	fmt.Println(count)
}
