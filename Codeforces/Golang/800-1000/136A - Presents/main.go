package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	result := make([]int, n+1)

	for i := 1; i <= n; i++ {
		var p int
		fmt.Scan(&p)
		result[p] = i
	}

	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", result[i])
	}
	fmt.Println()
}
