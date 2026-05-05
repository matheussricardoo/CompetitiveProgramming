package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	current := 0
	maxCapacity := 0

	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		current = current - a + b

		if current > maxCapacity {
			maxCapacity = current
		}
	}
	fmt.Println(maxCapacity)
}
